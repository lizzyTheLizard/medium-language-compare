class NotFoundException(Exception):
    def __init__(self, id):
        self.__id = id

    def getMessage(self):
        return "Could not find issue " + str(self.__id)        
