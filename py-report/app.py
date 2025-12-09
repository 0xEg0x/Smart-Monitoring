from flask import Flask, send_file, request, jsonify
import matplotlib
matplotlib.use('Agg')  # backend para gerar PNG sem display
import matplotlib.pyplot as plt
import io
import numpy as np
import pandas as pd
import datetime

app = Flask(__name__)

# Endpoint que gera gráfico PNG com base em dados enviados ou aleatórios
@app.route('/report', methods=['GET', 'POST'])
def report():
    # Aceita dados JSON tipo:
    # {"labels": [...], "values": [...]}
    data = request.get_json(silent=True)

    if data and 'values' in data:
        values = data['values']
        labels = data.get('labels', list(range(len(values))))
    else:
        labels = pd.date_range(
            end=datetime.datetime.now(),
            periods=10
        ).strftime('%H:%M:%S')

        values = np.random.randint(10, 100, size=10)

    # Geração do gráfico
    fig, ax = plt.subplots()
    ax.plot(labels, values)
    ax.set_title('Relatório gerado')
    ax.set_xlabel('Tempo')
    ax.set_ylabel('Valor')
    fig.autofmt_xdate()

    # Envia PNG como resposta
    buf = io.BytesIO()
    fig.savefig(buf, format='png', bbox_inches='tight')
    buf.seek(0)

    return send_file(buf, mimetype='image/png', download_name='report.png')


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
