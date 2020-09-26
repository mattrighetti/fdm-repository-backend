const db = require('./database.service');

module.exports = { getDropdownData }

async function getUniqueValuesOfField(field) {
    var array_data = []

    return db('models')
    .distinct(field)
    .then(array => {
        array.forEach(object => {
            if (object[field] != null) {
                array_data.push(object[field])
            }
        })

        return array_data
    })
}

/**
 * Removes every filter of type filer1/.../filterN 
 * and adds filter1,..., filterN to the array of filters 
 * if those values are not present
 * @param {Array<String>} filter_array
 */
async function removeDuplicateWithSeparator(filter_array) {
    var tmp = filter_array

    filter_array
    .filter(filter => filter.includes('/'))
    .forEach(filter => {
        let filter_index = tmp.indexOf(filter)
        if (filter_index > -1) {
            tmp.splice(filter_index, 1)
        }

        filter.split('/').forEach(split_value => {
            if (!tmp.includes(split_value)) {
                tmp.push(split_value)
            }
        })
    })
    
    return tmp
}

async function getDropdownData() {
    dict = {}
    promises = [
        getUniqueValuesOfField('cod').then(data => dict['cod'] = data),
        getUniqueValuesOfField('eis').then(data => removeDuplicateWithSeparator(data)).then(data => dict['eis'] = data),
        getUniqueValuesOfField('soa').then(data => removeDuplicateWithSeparator(data)).then(data => dict['soa'] = data),
        getUniqueValuesOfField('floodtypei').then(data => removeDuplicateWithSeparator(data)).then(data => dict['floodtypei'] = data),
        getUniqueValuesOfField('floodtypeii').then(data => removeDuplicateWithSeparator(data)).then(data => dict['floodtypeii'] = data),
        getUniqueValuesOfField('modeltypei').then(data => removeDuplicateWithSeparator(data)).then(data => dict['modeltypei'] = data),
        getUniqueValuesOfField('modeltypeii').then(data => removeDuplicateWithSeparator(data)).then(data => dict['modeltypeii'] = data),
        getUniqueValuesOfField('modeltypeiii').then(data => removeDuplicateWithSeparator(data)).then(data => dict['modeltypeiii'] = data)
    ]

    return Promise
    .all(promises)
    .then(() => dict )
}