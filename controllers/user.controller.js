const express = require('express');
const userServices = require('../services/user.services');
const router = express.Router();

module.exports = router;

// TODO changed endpoint name
router.post("/saveUserData", function (req, res) {
    let newUser = {
        name: req.body.name,
        surname: req.body.surname,
        phone: req.body.phone,
        email: req.body.email,
        affiliation: req.body.affiliation,
        comment: req.body.comment
    };

    userServices.saveUserData(newUser)
    .then(result => {
        res.json({
            success: result,
            message: 'Success'
        })
    })
});

router.post("/saveFile", async function (req, res) {
    userServices.saveFile(req.files)
    .then(result => {
        res
        .status(result ? 200 : 500)
        .send(result ? "File saved correctly" : "No file could be uploaded")
    }).catch(err => {
        res.status(500).send(err);
    })
});