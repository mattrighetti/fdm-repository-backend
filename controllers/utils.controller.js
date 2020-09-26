const express = require('express');
const router = express.Router()
const utilsServices = require('../services/utils.services');

module.exports = router

router.get('/filters', function(req, res) {
    utilsServices.getDropdownData().then(data => {
        res.json(data)
    })
})