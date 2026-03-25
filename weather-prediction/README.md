# Weather Prediction Sample

Blog article: [広告の「学習」って何してるの？](https://kkm-mako.com/blog/articles/ad-learning-machine-learning-explained/)

Predicts tomorrow's maximum temperature in Tokyo using historical weather data from [Open-Meteo API](https://open-meteo.com/).

## Quick Start

```bash
uv run predict_temperature.py
```

## What it does

1. Fetches 90 days of weather data for Tokyo (temperature, humidity, pressure, wind speed)
2. Trains a Random Forest model to predict next-day max temperature
3. Evaluates prediction accuracy on held-out test data

## Requirements

- Python 3.10+
- [uv](https://docs.astral.sh/uv/)
- No API key needed (Open-Meteo is free for non-commercial use)
