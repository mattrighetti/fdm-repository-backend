const db = require('./database.service');

module.exports = {
    saveUserData,
    saveFile
}

async function saveUserData(user) {
    var date = new Date();
    let month = date.getMonth() + 1;
    let year = date.getFullYear();

    user.month = month;
    user.year= year;

    db('user')
    .insert(user)
    .then(() => {
        return true;
    });
}

async function saveFile(files) {
    if (!files) {
        return false;
    } else {
        let file = files.file;
        const path = __dirname + '/other/uploaded_files/' + file.name;

        file.mv(path, (error) => {
            if (error) {
                console.error(error);
                throw error;
            }
            
            return true;
        });
    }
}