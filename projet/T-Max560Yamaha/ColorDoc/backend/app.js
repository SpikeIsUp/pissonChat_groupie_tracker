const express = require('express');

const app = express();
const port = 3000;

const cors = require('cors');
app.use(cors({
    origin: '*'
}))

const sneakerRouter = require('./router/sneaker');
app.use(sneakerRouter);

app.listen(port, () => console.log(`Server is running on port ${port}`));