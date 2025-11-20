import argparse
import json
import os
import re
from pathlib import Path
from typing import List, Dict, Any, Optional

import requests
from dotenv import load_dotenv


CHINESE_TRIPLE_REGEX = re.compile(r"head：(\{.*?\})；relationship：(\{.*?\})；tail：(\{.*?\})")


class ProviderConfig:
    def __init__(self, name: str, base_url: str, api_key_env: str, default_model: str):
        self.name = name
        self.base_url = base_url
        self.api_key_env = api_key_env
        self.default_model = default_model

# 模型提供商配置
PROVIDERS: Dict[str, ProviderConfig] = {
    # DeepSeek OpenAI-Compatible
    "deepseek": ProviderConfig(
        name="deepseek",
        base_url="https://api.deepseek.com",
        api_key_env="DEEPSEEK_API_KEY",
        default_model="deepseek-chat",
    ),
    # Qwen DashScope OpenAI-Compatible
    "qwen": ProviderConfig(
        name="qwen",
        base_url="https://dashscope.aliyuncs.com/compatible-mode/v1",
        api_key_env="DASHSCOPE_API_KEY",
        default_model="qwen2.5-72b-instruct",
    ),
    # Generic OpenAI-compatible forward proxy
    "forward": ProviderConfig(
        name="forward",
        base_url=os.environ.get("FORWARD_BASE_URL", "http://localhost:3000/v1"),
        api_key_env="FORWARD_API_KEY",
        default_model=os.environ.get("FORWARD_DEFAULT_MODEL", "gpt-4o-mini"),
    ),
}


def read_txt_text(txt_path: Path) -> str:
    """读取文本文件内容"""
    try:
        # 尝试不同的编码格式
        encodings = ['utf-8', 'gbk', 'gb2312', 'utf-16', 'latin-1']
        
        for encoding in encodings:
            try:
                return txt_path.read_text(encoding=encoding).strip()
            except UnicodeDecodeError:
                continue
        
        # 如果所有编码都失败，使用utf-8并忽略错误
        return txt_path.read_text(encoding='utf-8', errors='ignore').strip()
        
    except Exception as e:
        raise ValueError(f"无法读取文本文件 {txt_path}: {str(e)}")


def load_prompt(prompt_path: Path) -> str:
    return prompt_path.read_text(encoding="utf-8").strip()


def chunk_text(text: str, max_chars: int = 800) -> List[str]:
    """
    按自然段切分文本，并在必要时对超长段落按句子边界细分。

    策略：
    1) 先根据空行（连续换行）识别自然段。
    2) 每个自然段单独作为一个chunk，不合并多个段落。
    3) 若单个段落长度 > max_chars，则在该段内部按句子边界再切分。
    """
    text = text.strip()
    if not text:
        return []
    if len(text) <= max_chars:
        return [text]

    # 识别自然段：按空行（连续换行）分隔
    # 先标准化换行符
    normalized = text.replace('\r\n', '\n').replace('\r', '\n')
    
    # 按空行分割段落（匹配一个或多个连续换行，可能包含空白字符）
    # 这样可以把空行分隔的内容识别为不同段落
    raw_paragraphs = re.split(r'\n\s*\n+', normalized)
    
    # 清理每个段落：去除首尾空白，保留段落内部结构
    paragraphs = []
    for para in raw_paragraphs:
        cleaned = para.strip()
        if cleaned:  # 只保留非空段落
            paragraphs.append(cleaned)
    
    # 如果没有识别到段落（文本中无空行），则将整个文本作为一个段落
    if not paragraphs:
        paragraphs = [normalized.strip()]
    
    if not paragraphs:
        return [text]

    # 句子结束标点（用于细分超长段落）
    sentence_endings = ['。', '！', '？', '；', '…', '.', '!', '?', ';']

    def split_long_paragraph(para: str) -> List[str]:
        if len(para) <= max_chars:
            return [para]
        segments: List[str] = []
        start = 0
        while start < len(para):
            end = min(start + max_chars, len(para))
            if end >= len(para):
                segments.append(para[start:].strip())
                break
            # 在句子边界处后退查找
            split_pos = -1
            for i in range(end - 1, max(start, end - max_chars // 2) - 1, -1):
                if para[i] in sentence_endings:
                    split_pos = i + 1
                    break
            if split_pos == -1:
                # 退而求其次：按空白或标点
                for i in range(end - 1, max(start, end - max_chars // 2) - 1, -1):
                    if para[i] in [' ', '\n', '\t', '，', '、', ',', ':', '：']:
                        split_pos = i + 1
                        break
            if split_pos == -1:
                split_pos = end
            segments.append(para[start:split_pos].strip())
            start = split_pos
            while start < len(para) and para[start] in [' ', '\n', '\t']:
                start += 1
        return [seg for seg in segments if seg]

    # 每个自然段单独成为一个chunk，只在段落太长时按句子切分
    chunks: List[str] = []
    
    for para in paragraphs:
        if not para:
            continue
        
        if len(para) <= max_chars:
            # 段落长度合适，直接作为一个chunk
            chunks.append(para)
        else:
            # 段落太长，按句子边界切分
            segments = split_long_paragraph(para)
            for seg in segments:
                if len(seg) <= max_chars:
                    chunks.append(seg)
                else:
                    # 极端情况：句子仍超限，硬切
                    for i in range(0, len(seg), max_chars):
                        chunks.append(seg[i:i + max_chars])

    return [c for c in chunks if c]


def make_client(provider: str, model: Optional[str], base_url: Optional[str], api_key: Optional[str]):
    cfg = PROVIDERS[provider]
    resolved_base = base_url or cfg.base_url
    resolved_api_key = api_key or os.environ.get(cfg.api_key_env)
    if not resolved_api_key:
        raise RuntimeError(f"缺少API Key，请设置环境变量 {cfg.api_key_env} 或通过 --api-key 传入。")
    resolved_model = model or cfg.default_model
    return resolved_base, resolved_api_key, resolved_model


def call_llm(base_url: str, api_key: str, model: str, prompt_text: str, chunk: str, temperature: float = 0.0, top_p: float = 1.0) -> str:
    messages = [
        {"role": "system", "content": "你是一个严谨的中文信息抽取助手。"},
        {"role": "user", "content": f"{prompt_text}\n\n以下为需要抽取的原文：\n\n{chunk}"},
    ]
    
    headers = {
        "Authorization": f"Bearer {api_key}",
        "Content-Type": "application/json"
    }
    
    payload = {
        "model": model,
        "messages": messages,
        "temperature": temperature,
        "top_p": top_p
    }
    
    response = requests.post(f"{base_url}/chat/completions", headers=headers, json=payload)
    response.raise_for_status()
    
    result = response.json()
    return result["choices"][0]["message"]["content"].strip()


def parse_triples_from_text(text: str) -> List[Dict[str, Any]]:
    triples: List[Dict[str, Any]] = []
    
    # 首先尝试解析新的JSON格式
    try:
        # 查找JSON代码块
        json_start = text.find('```json')
        if json_start != -1:
            json_end = text.find('```', json_start + 7)
            if json_end != -1:
                json_content = text[json_start + 7:json_end].strip()
                data = json.loads(json_content)
                
                # 解析新的格式
                if "三元组列表" in data:
                    for triple in data["三元组列表"]:
                        if all(key in triple for key in ["head", "relationship", "tail"]):
                            triples.append(triple)
                    return triples
    except (json.JSONDecodeError, KeyError, TypeError):
        pass
    
    # 如果新格式解析失败，回退到原来的正则表达式解析
    for match in CHINESE_TRIPLE_REGEX.finditer(text.replace("\r", "")):
        head_str, rel_str, tail_str = match.groups()
        try:
            head_obj = json.loads(head_str)
            rel_obj = json.loads(rel_str)
            tail_obj = json.loads(tail_str)
            triples.append({"head": head_obj, "relationship": rel_obj, "tail": tail_obj})
        except json.JSONDecodeError:
            # 宽松兜底：尝试将中文引号替换
            fixed = head_str.replace("'", '"'), rel_str.replace("'", '"'), tail_str.replace("'", '"')
            try:
                head_obj = json.loads(fixed[0])
                rel_obj = json.loads(fixed[1])
                tail_obj = json.loads(fixed[2])
                triples.append({"head": head_obj, "relationship": rel_obj, "tail": tail_obj})
            except Exception:
                continue
    
    return triples


def save_outputs(raw_outputs: List[str], parsed: List[Dict[str, Any]], chunks: List[str], 
                 chunk_indices: List[int], out_txt: Path, out_jsonl: Path, chunks_txt: Path) -> None:
    """保存输出结果，包含三元组的chunk索引信息
    
    Args:
        raw_outputs: 原始LLM输出
        parsed: 解析后的三元组列表
        chunks: 文本分块列表
        chunk_indices: 每个三元组对应的chunk索引（用于溯源）
        out_txt: 输出txt文件路径
        out_jsonl: 输出jsonl文件路径
        chunks_txt: 输出chunks文件路径
    """
    out_txt.parent.mkdir(parents=True, exist_ok=True)
    out_jsonl.parent.mkdir(parents=True, exist_ok=True)
    chunks_txt.parent.mkdir(parents=True, exist_ok=True)

    # 保存原始输出
    out_txt.write_text("\n\n".join(raw_outputs), encoding="utf-8")
    
    # 为每个三元组添加chunk溯源信息
    enriched_parsed = []
    for i, triple in enumerate(parsed):
        enriched_triple = triple.copy()
        # 添加chunk索引信息
        if i < len(chunk_indices):
            chunk_idx = chunk_indices[i]
            enriched_triple["_chunk_index"] = chunk_idx
            # 确保索引有效
            if 0 <= chunk_idx < len(chunks):
                enriched_triple["_source_text"] = chunks[chunk_idx]
            else:
                enriched_triple["_source_text"] = f"[索引 {chunk_idx} 超出范围，共 {len(chunks)} 个chunks]"
        enriched_parsed.append(enriched_triple)
    
    with out_jsonl.open("w", encoding="utf-8") as f:
        json.dump(enriched_parsed, f, ensure_ascii=False, indent=2)
    
    # 保存分段文本
    chunk_content = []
    for i, chunk in enumerate(chunks, 1):
        chunk_content.append(f"=== 分段 {i} ===")
        chunk_content.append(f"长度: {len(chunk)} 字符")
        chunk_content.append(f"内容:\n{chunk}")
        chunk_content.append("")  # 空行分隔
    
    chunks_txt.write_text("\n".join(chunk_content), encoding="utf-8")


def process_single_txt(txt_path: Path, prompt_text: str, base_url: str, api_key: str, model: str, 
                       output_dir: Path, file_index: int, total_files: int) -> Dict[str, Any]:
    """处理单个txt文件"""
    print(f"[{file_index}/{total_files}] 处理文件: {txt_path.name}")
    
    # 生成输出文件名（基于原文件名）
    stem = txt_path.stem
    out_txt = output_dir / "txt" / f"{stem}_extractions.txt"
    out_jsonl = output_dir/ "jsonl" / f"{stem}_extractions.jsonl"
    chunks_txt = output_dir / "txt" / f"{stem}_chunks.txt"
    
    try:
        text = read_txt_text(txt_path)
        chunks = chunk_text(text)
        raw_outputs: List[str] = []
        parsed_triples: List[Dict[str, Any]] = []
        chunk_indices: List[int] = []  # 记录每个三元组来自哪个chunk

        for idx, chunk in enumerate(chunks, start=1):
            print(f"  处理分段 {idx}/{len(chunks)}，长度 {len(chunk)} 字符...")
            content = call_llm(base_url, api_key, model, prompt_text, chunk)
            raw_outputs.append(content)
            triples_from_chunk = parse_triples_from_text(content)
            parsed_triples.extend(triples_from_chunk)
            # 为每个三元组记录它来自哪个chunk（索引从0开始）
            chunk_indices.extend([idx - 1] * len(triples_from_chunk))

        save_outputs(raw_outputs, parsed_triples, chunks, chunk_indices, out_txt, out_jsonl, chunks_txt)
        
        print(f"  ✓ 完成。输出文件：{out_txt.name}, {out_jsonl.name}, {chunks_txt.name}")
        
        return {
            "file": txt_path.name,
            "status": "success",
            "chunks_processed": len(chunks),
            "triples_extracted": len(parsed_triples),
            "output_files": [str(out_txt), str(out_jsonl), str(chunks_txt)]
        }
        
    except Exception as e:
        print(f"  ✗ 处理失败: {str(e)}")
        return {
            "file": txt_path.name,
            "status": "error",
            "error": str(e)
        }


def process_folder(input_dir: Path, prompt_text: str, base_url: str, api_key: str, model: str, 
                  output_dir: Path) -> List[Dict[str, Any]]:
    """批量处理文件夹中的所有txt文件"""
    # 查找所有txt文件
    txt_files = list(input_dir.glob("*.txt"))
    if not txt_files:
        print(f"在目录 {input_dir} 中未找到任何.txt文件")
        return []
    
    print(f"找到 {len(txt_files)} 个txt文件，开始批量处理...")
    
    # 确保输出目录存在
    output_dir.mkdir(parents=True, exist_ok=True)

    results = []
    for i, txt_file in enumerate(txt_files, 1):

        # 如果目标目录下已有对应文件名的txt，则跳过
        if (output_dir / "txt" / f"{txt_file.stem}_extractions.txt").exists():
            print(f"目标目录下 {txt_file.stem}_extractions.txt 已存在，跳过处理")
            continue

        result = process_single_txt(txt_file, prompt_text, base_url, api_key, model, output_dir, i, len(txt_files))
        results.append(result)
    
    # 生成汇总报告
    summary_file = output_dir / "batch_processing_summary.json"
    with summary_file.open("w", encoding="utf-8") as f:
        json.dump(results, f, ensure_ascii=False, indent=2)
    
    print(f"\n批量处理完成！汇总报告已保存到: {summary_file}")
    
    # 打印统计信息
    successful = [r for r in results if r["status"] == "success"]
    failed = [r for r in results if r["status"] == "error"]
    total_triples = sum(r.get("triples_extracted", 0) for r in successful)
    
    print(f"成功处理: {len(successful)} 个文件")
    print(f"处理失败: {len(failed)} 个文件")
    print(f"总共抽取三元组: {total_triples} 个")
    
    return results


def main() -> None:
    parser = argparse.ArgumentParser(description="从TXT读取文本，使用Prompt调用大模型进行三元组抽取")
    
    # 输入参数组
    input_group = parser.add_mutually_exclusive_group(required=True)
    input_group.add_argument("--txt", type=Path, help="输入的单个txt文件路径")
    input_group.add_argument("--folder", type=Path, help="输入的txt文件所在文件夹路径")
    
    parser.add_argument("--prompt", type=Path, default=Path("triple_extraction_prompt.txt"), help="生成好的Prompt文件路径")
    parser.add_argument("--provider", choices=list(PROVIDERS.keys()), default="deepseek", help="模型提供方：deepseek/qwen/forward")
    parser.add_argument("--model", type=str, default=None, help="模型名称，留空则使用该provider的默认模型")
    parser.add_argument("--base-url", type=str, default=None, help="OpenAI兼容Base URL，留空则使用provider默认")
    parser.add_argument("--api-key", type=str, default=None, help="API Key，留空则读取各provider约定的环境变量")
    
    # 输出参数组
    parser.add_argument("--out", type=Path, default=Path("output/extractions.txt"), help="保存单个文件原始输出的txt路径")
    parser.add_argument("--out-jsonl", type=Path, default=Path("output/extractions.jsonl"), help="保存单个文件解析后三元组的jsonl路径")
    parser.add_argument("--chunks-txt", type=Path, default=Path("output/chunks.txt"), help="保存单个文件分段文本的txt路径")
    parser.add_argument("--output-dir", type=Path, default=Path("output"), help="批量处理时的输出目录")
    
    args = parser.parse_args()

    # 先加载 .env，以便后续读取环境变量
    load_dotenv(dotenv_path='.env')

    prompt_text = load_prompt(args.prompt)
    base_url, api_key, model = make_client(args.provider, args.model, args.base_url, args.api_key)

    if args.txt:
        # 处理单个文件
        text = read_txt_text(args.txt)
        chunks = chunk_text(text)
        raw_outputs: List[str] = []
        parsed_triples: List[Dict[str, Any]] = []
        chunk_indices: List[int] = []  # 记录每个三元组来自哪个chunk

        for idx, chunk in enumerate(chunks, start=1):
            print(f"处理分段 {idx}/{len(chunks)}，长度 {len(chunk)} 字符...")
            content = call_llm(base_url, api_key, model, prompt_text, chunk)
            raw_outputs.append(content)
            triples_from_chunk = parse_triples_from_text(content)
            parsed_triples.extend(triples_from_chunk)
            # 为每个三元组记录它来自哪个chunk（索引从0开始）
            chunk_indices.extend([idx - 1] * len(triples_from_chunk))

        save_outputs(raw_outputs, parsed_triples, chunks, chunk_indices, args.out, args.out_jsonl, args.chunks_txt)
        print(f"已完成。原始输出写入：{args.out}；解析后的JSONL写入：{args.out_jsonl}；分段文本写入：{args.chunks_txt}")
        
    elif args.folder:
        # 批量处理文件夹
        process_folder(args.folder, prompt_text, base_url, api_key, model, args.output_dir)


if __name__ == "__main__":
    main() 