def read_file(file_name):
    """
     This function will be used to read the text file in which the test will be written.

     @param  : file_name is the path used to read the file
     @return : str
    """
    file_content = ""
    try:
        with open(file_name, mode='r') as file:
            file_content = file_name.read()
    except Exception as e:
        print("Error while reading file : {}, got error : {}".format(file_name, e))
        raise e

    return file_content
