const db = require('./database.service');

module.exports = {
    getAllModels,
    countAllModels,
    getModel,
    getHeader,
    getExpression,
    getId,
    getModelInputs,
    getModelOutputs,
    getInfoOnCalibration,
    getInfoOnValidation,
    getTransferability,
    getBibliography,
    getComparison,
    getModelMarkdown
}

/**
 * Returns all models in the database with all their fields
 * @param {If `true` returns all models with just their filters values and no other fiels} justWithFilters 
 */
async function getAllModels(justWithFilters=false) {
    if (justWithFilters) {
        return db("models")
        .select('id', 'name', 'acronym', 'version', 'cod', 'soa', 'floodtypei', 'floodtypeii', 'modeltypei', 'modeltypeii', 'modeltypeiii', 'eis')
        .orderBy('name', 'asc')
        .orderBy('acronym', 'asc')
    } else {
        return db("models")
        .select()
    }
}

/**
 * Returns the number of models that are stored in the database
 */
async function countAllModels() {
    return db("models")
    .count()
}

/**
 * Returns model with selected identifier
 * @param {Model ID} id 
 */
async function getModel(id) {
    return db("models")
    .where('id', id)
    .select()
}

async function getModelMarkdown(id) {
    return db("model_filters")
    .where('id', id)
    .select('description')
}

async function getHeader(modelId) {
    return db("models")
    .where('id', modelId)
    .select('name', 'acronym', 'version', 'cod', 'soa', 'floodtypei', 'floodtypeii', 'modeltypei', 'modeltypeii', 'modeltypeiii', 'eis')
}

async function getExpression(modelId) {
    return db("models")
    .where('id', modelId)
    .select('expression')
}

async function getId(modelId) {
    return db("models")
    .where('id', modelId)
    .select('name', 'acronym', 'version', 'yearoflastupdate', 'authors', 'linkedmodels', 'expression')
}

async function getModelInputs(modelId) {
    return db("models")
    .where('id', modelId)
    .select('modelinputs')
}

async function getModelOutputs(modelId) {
    return db("models")
    .where('id', modelId)
    .select('modeloutputs')
}

async function getInfoOnCalibration(modelId) {
    return db("models")
    .where('id', modelId)
    .select('infooncalibration')
}

async function getInfoOnValidation(modelId) {
    return db("models")
    .where('id', modelId)
    .select('infoonvalidation')
}

async function getTransferability(modelId) {
    return db("models")
    .where('id', modelId)
    .select('transferability')
}

async function getBibliography(modelId) {
    return db("models")
    .where('id', modelId)
    .select('bibliography')
}

async function getComparison(modelId) {
    return db("models")
    .where('id', modelId)
    .select('Comparison')
}