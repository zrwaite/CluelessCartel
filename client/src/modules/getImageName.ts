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
            filename = 'rv.png'
            break
        case "Rocks":
            filename = 'rock.png'
            break
        case "Trees":
            filename = 'tree.png'
            break
        case "Garbage":
            filename = 'garbage.png'
            break
        case "Buried Storage":
            filename = 'buriedStorage.png'
            break
        case "Armory":
            filename = 'armory.png'
            break
        case "Farm":
            filename = 'farm.jpg'
            break
        default: 
            found = false;
    }
    return {filename, found}
}