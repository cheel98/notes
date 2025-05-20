## 利用SSH隧道连接局域网主机

两台只有局域网的机器A和B，一台有公网IP的机器C，A和B不在同一局域网，如何在A上ssh连接B

### 建立反向隧道

```bash
ssh -R 2222:localhost:22 user@C_PUBLIC_IP
```

### 在A上通过C连接到B

```bash
ssh -J user@C_PUBLIC_IP:22 user@localhost -p 2222
```

### 在B上连接C

```bash
ssh root@localhost -p 2222
```

