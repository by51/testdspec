# Spec: string-utils

## ADDED Requirements

### Requirement: 按空格拆分字符串
系统应当提供一个函数，接受一个字符串参数，按空格将其拆分为字符串数组并返回。

#### Scenario: 普通字符串拆分
- **WHEN** 输入为 `"hello world foo"`
- **THEN** 返回 `["hello", "world", "foo"]`

#### Scenario: 连续空格视为单一分隔符
- **WHEN** 输入为 `"hello  world"`（两个空格）
- **THEN** 返回 `["hello", "world"]`，不包含空元素

#### Scenario: 首尾有空格
- **WHEN** 输入为 `"  hello world  "`
- **THEN** 返回 `["hello", "world"]`，不包含首尾空元素

#### Scenario: 空字符串输入
- **WHEN** 输入为 `""`
- **THEN** 返回空数组 `[]`

#### Scenario: 仅含空格的字符串
- **WHEN** 输入为 `"   "`
- **THEN** 返回空数组 `[]`

#### Scenario: 单词无空格
- **WHEN** 输入为 `"hello"`
- **THEN** 返回 `["hello"]`
