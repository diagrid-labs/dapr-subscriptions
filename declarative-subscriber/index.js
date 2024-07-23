import express from 'express';
import bodyParser from 'body-parser';

const APP_PORT = process.env.APP_PORT ?? '5004';

const app = express();
app.use(bodyParser.json({ type: 'application/*+json' }));

// Dapr subscription routes orders topic to this route
app.post('/orders', (req, res) => {
    console.log("Declarative Subscriber received:", req.body.data);
    res.sendStatus(200);
});

app.listen(APP_PORT);