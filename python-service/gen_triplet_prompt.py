import argparse
import json
from collections import defaultdict
from pathlib import Path
from typing import Dict, Set, List, Tuple

import pandas as pd


def load_schema(schema_path: Path) -> dict:
    with schema_path.open("r", encoding="utf-8") as f:
        return json.load(f)


def collect_types(schema: dict) -> Tuple[Dict[str, Set[str]], List[str]]:
    """Collect entity types (with includes) and relationship types from schema."""
    entity_type_to_children: Dict[str, Set[str]] = defaultdict(set)
    relationship_types: Dict[str, Set[str]] = defaultdict(set)
    triplet_types: Set[str] = set()


    for triple in schema.get("triples", []):
        triplet_str=""

        relationship=triple.get("relationship", {})

        head = triple.get("head", {})
        tail = triple.get("tail", {})

        relationship_type = relationship.get("type", "").strip()
        head_type = head.get("type", "").strip()
        tail_type = tail.get("type", "").strip()

        triplet_str+=head_type+'->'+triple.get("relationship", "").get("type", "")+'->'+tail_type
        triplet_types.add(triplet_str)          
        
        
        # Merge includes for entity types
        if head_type:
            for child in head.get("includes", []) or []:
                if child:
                    entity_type_to_children[head_type].add(str(child).strip())
            # Ensure the type key exists even if no includes
            entity_type_to_children.setdefault(head_type, set())
        if tail_type:
            for child in tail.get("includes", []) or []:
                if child:
                    entity_type_to_children[tail_type].add(str(child).strip())
            entity_type_to_children.setdefault(tail_type, set())
        
        if relationship_type:
            for child in relationship.get("includes", []):
                if child:
                    relationship_types[relationship_type].add(str(child).strip())
            relationship_types.setdefault(relationship_type, set())

    # Remove empty relationship names if any
    # relationship_types = {r.get("type", ""):r.get("includes", []) for r in relationship_types}
    return entity_type_to_children, relationship_types,triplet_types

def collect_triplet_types(schema: dict) -> Tuple[Dict[str, Set[str]], List[str]]:
    entity_type_to_children: Dict[str, Set[str]] = defaultdict(set)
    relationship_types: Set[str] = set()


    for triple in schema.get("triples", []):
        relationship_types.add(triple.get("relationship", "").strip())
        head = triple.get("head", {})
        tail = triple.get("tail", {})


def build_prompt(entity_type_to_children: Dict[str, Set[str]], relationship_types: Dict[str, Set[str]],triplet_types: Set[str]) -> str:
    lines: List[str] = []
    lines.append("你是东南大学建筑学院的一个行政人员。输入的文本代表你院的一些实践活动记录，请按照步骤从输入文本中抽取三元组。")
    lines.append("")
    lines.append("步骤1：先通篇理解，抽取文本中唯一的实践活动")
    lines.append("步骤2：围绕实践活动，抽取以下符合下述知识图谱模式的实体和三元组")
    lines.append("步骤3：抽取参与实践活动的人物，以及在实践活动中的实践行为")
    lines.append("步骤4：然后针对实践活动结合原文总结实践成果，并且总结实践成果依托的专业技能")

    lines.append("")
    lines.append("=== 重要：输出格式要求 ===")
    lines.append("你必须严格按照以下JSON格式输出，不要添加任何其他文字、说明或格式：")
    lines.append("")
    lines.append("```json")
    lines.append("{")
    lines.append('  "实践活动": "实践活动名称",')
    lines.append('  "三元组列表": [')
    lines.append('    {')
    lines.append('      "head": {"label": "实体表面词", "type": "实体类型"},')
    lines.append('      "relationship": {"label": "关系表面词", "type": "关系类型"},')
    lines.append('      "tail": {"label": "实体表面词", "type": "实体类型"}')
    lines.append('    }')
    lines.append('  ]')
    lines.append("}")
    lines.append("```")
    lines.append("")
    lines.append("=== 输出示例（仅格式参考） ===")
    lines.append("```json")
    lines.append("{")
    lines.append('  "实践活动": "[从原文中提取的实践活动名称]",')
    lines.append('  "三元组列表": [')
    lines.append('    {')
    lines.append('      "head": {"label": "[原文中的实体表面词]", "type": "[实体类型]"},')
    lines.append('      "relationship": {"label": "[原文中的关系表面词]", "type": "[关系类型]"},')
    lines.append('      "tail": {"label": "[原文中的实体表面词]", "type": "[实体类型]"}')
    lines.append('    }')
    lines.append('  ]')
    lines.append("}")
    lines.append("```")
    lines.append("")
    lines.append("**注意：请从您收到的原文中抽取三元组，不要使用示例中的具体内容！**")
    lines.append("")
    lines.append("=== 抽取规则 ===")
    lines.append("- 必填字段：`label` 为原文中的表面词；`type` 必须从下方枚举中选择。")
    lines.append("- 如 tail 的表面词来自某类型的`细分`（includes），则 `type` 写该父类型名称，`label` 写细分项文本（如果不在其中，也可以自行总结细分项）。")
    lines.append("- 关系的 `type` 也必须来自下方关系类型枚举，`label` 写原文触发词（可与 type 相同或为同义表达）。")
    lines.append("- 三元组的类型必须来自下方三元组类型枚举，不要出现其他的类型，并且不要头尾颠倒")
    lines.append("- 当遇见并列的'、'、'，'等表达，将并列的实体分别抽取")
    lines.append("- 严格按照上述JSON格式输出，不要添加任何其他内容")
    lines.append("- 确保JSON格式完全正确，可以被直接解析")
    lines.append("")

    lines.append("实体类型与细分（type -> includes）：")
    for etype in sorted(entity_type_to_children.keys()):
        children = sorted([c for c in entity_type_to_children[etype] if c])
        if children:
            lines.append(f"- {etype}：{', '.join(children)}")
        else:
            lines.append(f"- {etype}：<无细分>")
    lines.append("")

    lines.append("关系类型与细分（type -> includes）：")
    if relationship_types:
        for rtype in sorted(relationship_types):
            children = sorted([c for c in relationship_types[rtype] if c])
            if children:
                lines.append(f"- {rtype}：{', '.join(children)}")
            else:
                lines.append(f"- {rtype}：<无细分>")
    else:
        lines.append("- <未在schema中定义>")
    lines.append("")
    
    # 重新组织三元组类型，按头实体分类
    lines.append("三元组类型（type）：")
    if triplet_types:
        # 按头实体分组
        head_entity_groups = defaultdict(list)
        for triplet in triplet_types:
            parts = triplet.split('->')
            if len(parts) == 3:
                head_entity = parts[0]
                head_entity_groups[head_entity].append(triplet)
        
        # 按头实体分组显示
        for head_entity in sorted(head_entity_groups.keys()):
            lines.append(f"\n{head_entity}相关：")
            for triplet in sorted(head_entity_groups[head_entity]):
                lines.append(f"  {triplet}")
    else:
        lines.append("- <未在schema中定义>")
    lines.append("")
    
    lines.append("=== 再次强调 ===")
    lines.append("请严格按照JSON格式输出，不要添加任何其他文字、说明或格式。")
    lines.append("输出必须是有效的JSON，可以被直接解析。")
    lines.append("")

    lines.append("抽取样例：")
    lines.append("原文：")
    docx_path=Path("/Users/hsy/Documents/WaWa/Proj/HaborKG/archikg_2/【回顾】南师附小弹性离校 _ 放学别走！一起 识天文知地理！.txt")
    docx_text=read_txt_text(docx_path)
    lines.append(docx_text)
    lines.append("")
    lines.append("抽取结果：")
    xlsx_path=Path("/Users/hsy/Documents/WaWa/Proj/HaborKG/archikg_2/【回顾】南师附小弹性离校 _ 放学别走！一起 “识天文知地理”！_extractions.xlsx")
    xlsx_text=read_xlsx_text(xlsx_path)
    lines.append(xlsx_text)
    lines.append("")



    return "\n".join(lines)
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

def read_xlsx_text(xlsx_path: Path) -> str:
    df = pd.read_excel(xlsx_path)
    # 删除空行
    df = df.dropna(how='all')
    return df.to_string()


def main() -> None:
    parser = argparse.ArgumentParser(description="根据知识图谱schema生成三元组抽取Prompt")
    parser.add_argument("--schema", type=Path, default=Path("knowledge_graph_schema.json"), help="schema JSON 文件路径")
    parser.add_argument("--out", type=Path, default=Path("output/triple_extraction_prompt.txt"), help="输出的prompt文件路径")
    args = parser.parse_args()

    schema = load_schema(args.schema)
    entity_type_to_children, relationship_types,triplet_types = collect_types(schema)


    prompt_text = build_prompt(entity_type_to_children, relationship_types,triplet_types)

    # Ensure parent dir exists
    args.out.parent.mkdir(parents=True, exist_ok=True)
    args.out.write_text(prompt_text, encoding="utf-8")

    # 同时打印到标准输出，便于直接查看
    print(prompt_text)


if __name__ == "__main__":
    main() 