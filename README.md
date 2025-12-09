📊 Monitoramento com Prometheus + Grafana + Go API + Python Exporter

Este projeto demonstra como integrar uma API em Go, um exporter em Python, o Prometheus e o Grafana para criação de um pipeline completo de monitoramento com métricas personalizadas e dashboards visuais.

🚀 Tecnologias utilizadas

Go + Echo + Gorm — API principal

Python 3 + prometheus_client — Exporter customizado

Prometheus — Coleta de métricas dos serviços

Grafana — Visualização dos gráficos

Docker & Docker Compose — Orquestração

SQLite — Banco utilizado pela API em Go

🧱 Arquitetura
                ┌──────────────────────────────┐
                │            GRAFANA           │
                │    (Visualização de dados)   │
                └───────────────▲──────────────┘
                                │
                                │ consulta
                                │
                ┌───────────────┴──────────────┐
                │           PROMETHEUS          │
                │ (Coleta métricas de Go e Py)  │
                └───────────────▲──────────────┘
                    scrape      │         scrape
                                │
      ┌─────────────────────────┘───────────────┐
      │                                         │
┌───────────────┐                     ┌─────────────────┐
│ Go API (Echo) │   /metrics endpoint │ Python Exporter │
│ + Student API │────────────────────▶│ custom metrics  │
└───────────────┘                     └─────────────────┘

📦 Estrutura do projeto
/project-root
  /api
    main.go
    go.mod
  /python-exporter
    exporter.py
  /prometheus
    prometheus.yml
  docker-compose.yml
  README.md

⚙️ Como rodar o projeto
1️⃣ Clonar o repositório
git clone https://github.com/seuuser/seuprojeto
cd seuprojeto

2️⃣ Subir todos os serviços
docker compose up -d --build

3️⃣ Acessar as ferramentas
Serviço	URL
Go API	http://localhost:8080

Métricas Go	http://localhost:8080/metrics

Exporter Py	http://localhost:8000/metrics

Prometheus	http://localhost:9090

Grafana	http://localhost:3000

Login padrão Grafana:
user: admin
pass: admin

🐍 Python Exporter (gráficos + métricas)

Você pode colocar QUALQUER código Python e gerar métricas a partir do que quiser:

✔ estatísticas de banco
✔ temperatura da CPU
✔ dados de sensores
✔ ping, latência, upload/download
✔ geração de gráficos (matplotlib, seaborn, plotly)
✔ análise de arquivos

Um exemplo simples do exporter:

from prometheus_client import start_http_server, Gauge
import time, random

cpu_temp = Gauge('python_cpu_temp', 'Temperatura simulada da CPU')

if __name__ == '__main__':
    start_http_server(8000)
    while True:
        cpu_temp.set(random.uniform(40, 80))
        time.sleep(2)


Simples, poderoso e 100% integrável ao Prometheus.

📈 Criando dashboards no Grafana

Abra o Grafana

Adicione datasource → Prometheus

Use queries como:

python_cpu_temp
go_api_requests_total


Monte gráficos, gauges, heatmaps — tudo em modo visual.

📡 Prometheus Config (scrapes)

Arquivo prometheus.yml:

scrape_configs:
  - job_name: 'go-api'
    static_configs:
      - targets: ['api:8080']

  - job_name: 'python-exporter'
    static_configs:
      - targets: ['python-exporter:8000']

🔥 Comandos úteis

Logs de um serviço:

docker compose logs -f api


Rebuild:

docker compose up -d --build


Parar:

docker compose down