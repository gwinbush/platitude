from flask import Flask, jsonify
import random
import os
import time

app = Flask(__name__)

platitudes = [
    "There's only one obvious way to do it... until you find three more on Stack Overflow.",
    "Indentation isn't just style, it's sacred.",
    "When in doubt, import this.",
    "Explicit is better than implicit... unless it makes the code shorter.",
    "Don't reinvent the wheel, just pip install it.",
]

# Initialize secure random with current time as seed
secure_random = random.SystemRandom()


@app.route("/python/platitude")
def get_platitude():
    # Reseed with current time to ensure randomness
    random.seed(time.time())

    selected_platitude = secure_random.choice(platitudes)

    return jsonify(
        {
            "message": selected_platitude,
        }
    )


@app.route("/health")
def health():
    return jsonify({"status": "healthy"})


if __name__ == "__main__":
    port = int(os.environ.get("PORT", 80))
    app.run(host="0.0.0.0", port=port)
