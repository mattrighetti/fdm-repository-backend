const db = require('./database.service');

module.exports = { getDataChart }

async function getDataChart() {
    return db('models')
    .select('cod', 'name')
}