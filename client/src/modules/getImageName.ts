export const getImageName = (filename: string) => {
    return '../../assets/' + filename;
}

export const getStructureImage = (structureName: string): {filename: string, found: boolean} => {
    let filename = ''
    let found = true
    switch (structureName) {
        case "Shed":
            filename = 'shed.png'
            break
        case "Storage Unit":
            filename = 'storage.jpg'
            break
        case "RV": 
            filename = 'truck.png'
            break
        default: 
            found = false;
    }
    return {filename, found}
}