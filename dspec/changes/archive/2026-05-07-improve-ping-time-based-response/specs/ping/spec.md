## MODIFIED Requirements

### Requirement: ping 接口响应

#### Scenario: 工作时段内调用
- **WHEN** 请求 `/ping`，且服务器本地时间在 10:00（含）至 18:00（不含）之间
- **THEN** 响应体为 `pong`，状态码 200

#### Scenario: 非工作时段调用
- **WHEN** 请求 `/ping`，且服务器本地时间早于 10:00 或不早于 18:00
- **THEN** 响应体为 `pang`，状态码 200
