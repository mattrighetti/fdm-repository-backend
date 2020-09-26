const db = require('./database.service');

async function getAllMissingModels() {
    return db("missing_models")
    .select('id', 'name', 'biblio')
}

async function countMissingModels() {
    return db("missing_models")
    .count()
}

module.exports = { 
    getAllMissingModels,
    countMissingModels
}