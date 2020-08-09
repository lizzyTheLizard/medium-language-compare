class Issue(object):
    def __init__(self, id, name, description):
        self.__id = id
        self.__name = name
        self.__description = description
    
    def getId(self):
        return self.__id
        
    def getName(self):
        return self.__name

    def getDescription(self):
        return self.__description

    def update(self, newName, newDescription):
        if(newName == None or newName == ""):
            newName = self.__name
        if(newDescription == None or newDescription == ""):
            newDescription = self.__description
        return Issue(self.__id, newName, newDescription)

