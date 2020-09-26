const express = require('express');
const modelsServices = require('../services/models.services');
const router = express.Router()

module.exports = router;

router.get("/", function (req, res) {
    modelsServices.getAllModels(justWithFilters=true)
    .then(result => {
        res.json(result);
    });
});

router.get("/all", function (req, res) {
    modelsServices.getAllModels()
    .then(result => {
        res.json(result);
    });
});

router.get("/count", function (req, res) {
    modelsServices.countAllModels()
    .then(result => {
        res.json(result);
    });
});

router.get("/:modelId", function (req, res) {
    let idm = parseInt(req.params.modelId);
    modelsServices.getModel(idm)
    .then(result => {
        res.json(result);
    });
});

router.get("/:modelId/markdown", function (req, res) {
    let idm = parseInt(req.params.modelId);
    modelsServices.getModelMarkdown(idm)
    .then(result => {
        res.json(result[0]);
    });
});

router.get("/:modelId/header", function (req, res) {
    let idm = parseInt(req.params.modelId);
    modelsServices.getHeader(idm)
    .then(result => {
        res.json(result);
    });
});

router.get("/:modelId/expression", function (req, res) {
    let idm = parseInt(req.params.modelId);
    modelsServices.getExpression(idm)
    .then(result => {
        // TODO this could be unwrapped frontend
        res.json(result[0].expression);
    });
});

router.get("/:modelId/id", function (req, res) {
    let idm = parseInt(req.params.modelId);
    modelsServices.getId(idm)
    .then(result => {
        res.json(result);
    });
});

router.get("/:modelId/modelinputs", function (req, res) {
    let idm = parseInt(req.params.modelId);
    modelsServices.getModelInputs(idm)
    .then(result => {
        res.json(result);
    });
});

router.get("/:modelId/modeloutputs", function (req, res) {
    let idm = parseInt(req.params.modelId);
    modelsServices.getModelOutputs(idm)
    .then(result => {
        res.json(result);
    });
});

router.get("/:modelId/infooncalibration", function (req, res) {
    let idm = parseInt(req.params.modelId);
    modelsServices.getInfoOnCalibration(idm)
    .then(result => {
        res.json(result);
    });
});

router.get("/:modelId/infoonvalidation", function (req, res) {
    let idm = parseInt(req.params.modelId);
    modelsServices.getInfoOnValidation(idm)
    .then(result => {
        res.json(result);
    });
});

router.get("/:modelId/transferability", function (req, res) {
    let idm = parseInt(req.params.modelId);
    modelsServices.getTransferability(idm)
    .then(result => {
        res.json(result);
    });
});

router.get("/:modelId/bibliography", function (req, res) {
    let idm = parseInt(req.params.modelId);
    modelsServices.getBibliography(idm)
    .then(result => {
        res.json(result);
    });
});

router.get("/:modelId/comparison", function (req, res) {
    let idm = parseInt(req.params.modelId);
    modelsServices.getComparison(idm)
    .then(result => {
        res.json(result);
    });
});