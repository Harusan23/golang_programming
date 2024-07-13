package main

func minOperations(logs []string) int {
	currentPath := []string {}
	for i:=0; i<len(logs); i++{
		if (len(currentPath) == 0 && logs[i] == "../") || logs[i] == "./" {
            continue
        } else if len(currentPath) > 0 && logs[i] == "../" {
            currentPath = currentPath[:len(currentPath)-1]
        } else {
            currentPath = append(currentPath, logs[i])
        }
	}
    return len(currentPath)
}