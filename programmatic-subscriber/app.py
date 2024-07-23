from flask import Flask, request, jsonify
from cloudevents.http import from_http
import json
import os
import logging

app = Flask(__name__)

app_port = os.getenv('APP_PORT', '5002')

# Disable werkzeug request logging
log = logging.getLogger('werkzeug')
log.setLevel(logging.ERROR)

# Register Dapr pub/sub subscriptions
@app.route('/dapr/subscribe', methods=['GET'])
def subscribe():
    subscriptions = [{
        'pubsubname': 'pubsub',
        'topic': 'orders',
        'route': 'orders'
    }]
    print('Dapr pub/sub is subscribed to: ' + json.dumps(subscriptions))
    return jsonify(subscriptions)


# Dapr subscription in /dapr/subscribe sets up this route
@app.route('/orders', methods=['POST'])
def orders_subscriber():
    event = from_http(request.headers, request.get_data())
    print('Programmatic Subscriber received: %s' % event.data['orderId'], flush=True)
    return json.dumps({'success': True}), 200, {
        'ContentType': 'application/json'}


app.run(port=app_port)
