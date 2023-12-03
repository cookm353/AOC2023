#!/usr/bin/python3

digitMap = {
    "zero": "z0o",
    "one": "o1e",
    "two": "t20",
    "three": "t3e",
    "four": "4",
    "five": "5e",
    "six": "6",
    "seven": "7n",
    "eight": "e8t",
    "nine": "n9e"
}

def readFile(path: str) -> list[str]:
    """Read input from file into array of strings

    Args:
        path (str): Relative path for file to read

    Returns:
        list[str]: Input from file converted to an array of strings
    """
    with open(path, "r") as file:
        lines = file.readlines()

    return lines

def replaceNumbers(line: str) -> str:
    """Replace numbers spelled out with digits

    Args:
        line (str): Line from input

    Returns:
        str: Line with replaced numbers
    """
    
    for k, v in digitMap.items():
        if k in line:
            line = line.replace(k, v)
            
    return line[:-1]

def getCalibrationValue(line: str) -> int:
    """_summary_

    Args:
        line (str): _description_

    Returns:
        int: _description_
    """

    # print(line, end=" = ")
    line = list(filter(lambda char: char.isnumeric(), line))
    line = line[0] + line[-1]
    return int(line)
            

def main():
    lines = readFile("Day1/input1.txt")
    newLines = []
    sumCalibrationValues = 0
    
    for line in lines:
        newLines.append(replaceNumbers(line))
        
    for line in newLines:
        sumCalibrationValues += getCalibrationValue(line)
        
    print(sumCalibrationValues)
    

if __name__ == "__main__":
    main()