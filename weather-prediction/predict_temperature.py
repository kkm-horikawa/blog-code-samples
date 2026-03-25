"""
天気予報の予測サンプル — 翌日の最高気温を機械学習で予測する

ブログ記事「広告の『学習』って何してるの？」のサンプルコード
https://kkm-mako.com/blog/articles/ad-learning-machine-learning-explained/

使い方:
  pip install requests scikit-learn pandas
  python predict_temperature.py
"""

import json
import urllib.request
from datetime import datetime, timedelta

import pandas as pd
from sklearn.ensemble import RandomForestRegressor
from sklearn.model_selection import train_test_split
from sklearn.metrics import mean_absolute_error


def fetch_weather_data(latitude: float, longitude: float, days: int = 90) -> pd.DataFrame:
    """Open-Meteo API から過去の気象データを取得する"""
    end_date = datetime.now() - timedelta(days=1)
    start_date = end_date - timedelta(days=days)

    url = (
        f"https://api.open-meteo.com/v1/forecast"
        f"?latitude={latitude}&longitude={longitude}"
        f"&daily=temperature_2m_max,temperature_2m_min"
        f"&hourly=temperature_2m,relative_humidity_2m,pressure_msl,wind_speed_10m"
        f"&start_date={start_date:%Y-%m-%d}"
        f"&end_date={end_date:%Y-%m-%d}"
        f"&timezone=Asia%2FTokyo"
    )

    with urllib.request.urlopen(url) as resp:
        data = json.loads(resp.read())

    # 時間別データを日別に集約
    hourly = pd.DataFrame(data["hourly"])
    hourly["date"] = pd.to_datetime(hourly["time"]).dt.date
    daily_features = hourly.groupby("date").agg(
        avg_temp=("temperature_2m", "mean"),
        avg_humidity=("relative_humidity_2m", "mean"),
        avg_pressure=("pressure_msl", "mean"),
        avg_wind_speed=("wind_speed_10m", "mean"),
    ).reset_index()

    # 日別の最高気温
    daily = pd.DataFrame({
        "date": pd.to_datetime(data["daily"]["time"]).date,
        "max_temp": data["daily"]["temperature_2m_max"],
    })

    df = daily_features.merge(daily, on="date")
    return df


def prepare_dataset(df: pd.DataFrame) -> tuple:
    """予測の手がかり（特徴量）と正解データを作成する"""
    # 翌日の最高気温を「正解」として追加
    df["next_day_max_temp"] = df["max_temp"].shift(-1)
    df = df.dropna()

    # 過去の気象データから「予測の手がかり」を作る
    features = df[["avg_temp", "avg_humidity", "avg_pressure", "avg_wind_speed"]]

    # 「正解」= 翌日の最高気温
    labels = df["next_day_max_temp"]

    return features, labels


def main():
    print("=" * 60)
    print("天気予報の予測サンプル — 翌日の最高気温を予測する")
    print("=" * 60)

    # 東京の緯度経度
    TOKYO_LAT, TOKYO_LON = 35.6762, 139.6503

    print("\n[1] Open-Meteo API から東京の過去90日分の気象データを取得中...")
    df = fetch_weather_data(TOKYO_LAT, TOKYO_LON, days=90)
    print(f"    取得完了: {len(df)}日分のデータ")

    print("\n[2] 予測の手がかり（特徴量）と正解データを準備中...")
    features, labels = prepare_dataset(df)
    print(f"    手がかり: 平均気温, 平均湿度, 平均気圧, 平均風速")
    print(f"    正解:     翌日の最高気温")

    # 学習用データとテスト用データに分割
    X_train, X_test, y_train, y_test = train_test_split(
        features, labels, test_size=0.2, random_state=42
    )
    print(f"    学習用: {len(X_train)}日分 / テスト用: {len(X_test)}日分")

    print("\n[3] モデルを作って学習させる...")
    model = RandomForestRegressor(n_estimators=100, random_state=42)
    model.fit(X_train, y_train)  # ← ここが「学習」
    print("    学習完了")

    print("\n[4] テストデータで予測する...")
    predictions = model.predict(X_test)  # ← ここが「予測（推論）」

    mae = mean_absolute_error(y_test, predictions)
    print(f"    平均誤差: {mae:.1f}°C")

    print("\n[5] 予測結果の一部:")
    print(f"    {'予測':>8s}  {'実際':>8s}  {'誤差':>8s}")
    print(f"    {'-' * 30}")
    for pred, actual in list(zip(predictions, y_test))[:10]:
        diff = pred - actual
        print(f"    {pred:>7.1f}°C  {actual:>7.1f}°C  {diff:>+7.1f}°C")

    print(f"\n{'=' * 60}")
    print("まとめ:")
    print(f"  - 過去90日分の気象データで学習し、翌日の最高気温を予測")
    print(f"  - 平均誤差 {mae:.1f}°C で予測できた")
    print(f"  - 手がかりが多いほど、データが多いほど、精度は上がる")
    print(f"{'=' * 60}")


if __name__ == "__main__":
    main()
