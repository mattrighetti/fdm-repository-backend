const express = require("express");
const app = express();
const bodyParser = require("body-parser");
const process = require("process");
const _ = require("lodash");
const fileupload = require('express-fileupload')
const compression = require('compression')
const morgan = require('morgan');
const cors = require('cors');

let serverPort = process.env.PORT || 5000;

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({
    extended: true
}));
app.use(compression())
app.use(fileupload())
app.use(morgan('combined'));
app.use(cors());

app.use('/api/models',          require('./controllers/models.controller'));
app.use('/api/missing_models',  require('./controllers/missing_models.controller'));
app.use('/api/user',            require('./controllers/user.controller'));
app.use('/api/data',            require('./controllers/data.controller'));
app.use('/api/filters',         require('./controllers/filters.controller'));
app.use('/api/utils',           require('./controllers/utils.controller'));

app.set("port", serverPort);

/* Start the server on port 5000 */
app.listen(serverPort, function () {
    console.log(`Your app is ready at port ${serverPort}`);
});