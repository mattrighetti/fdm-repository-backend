const knex = require('knex'); 
const db_dev_config = require('../config/database_dev.json');
const db_prod_config = require('../config/database_prod.json');

let db;

function init() {
    console.log("Init database");
    if (process.env.TEST) {
        db = knex(db_dev_config);
    } else {
        db = knex(db_prod_config);
    }
}

init();

module.exports = db;