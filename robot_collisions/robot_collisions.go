package main
import (
	"sort"
	"fmt"
)

type RobotData struct {
	Position  int
	Health    int
	Direction string
	id int
}

func isNotEmpty(moves []int) bool {
	if len(moves) == 0 {
		return false
	}
	return true
}
func top(moves []int) int {
	lastPos := len(moves) - 1
	return moves[lastPos]
}
func pop(moves []int) []int {
	lastPos := len(moves) - 1
	temp := moves[:lastPos]
	return temp
}

func leftedHealths(data []RobotData, pos1, pos2 int) []RobotData {
	if data[pos1].Health > data[pos2].Health {
		data[pos1].Health = data[pos1].Health - 1
		data[pos2].Health = 0
	} else if data[pos1].Health < data[pos2].Health {
		data[pos2].Health = data[pos2].Health - 1
		data[pos1].Health = 0
	} else {
		data[pos1].Health = 0
		data[pos2].Health = 0
	}
	return data
}

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	rightMoves, leftMoves := []int{}, []int{} // use as the stacks
	var leftBotI, rightBotI int
	result := []int{}
	len_healths := len(healths)
	robot_data_list := []RobotData{}
	// prepare the data for computing
	for i := 0; i < len_healths; i++ { // add the data to the new data type's list
		robot_data_list = append(robot_data_list, RobotData{positions[i], healths[i], string(directions[i]), i})
	}
	sort.SliceStable(robot_data_list, func(k, j int) bool { // sort the robot data list by position
		return robot_data_list[k].Position < robot_data_list[j].Position
	})
	// fmt.Printf("start data: %v\n", robot_data_list)
	// compute the answer
	for index, data := range robot_data_list {
		// fmt.Printf("#%v round, ", index)
		if data.Health > 0 {
			if data.Direction == "R" && isNotEmpty(leftMoves) {
				leftBotI = top(leftMoves)
				if data.Position < robot_data_list[leftBotI].Position {
					for true {
						if data.Position > robot_data_list[leftBotI].Position || !isNotEmpty(leftMoves) || robot_data_list[index].Health == 0 {
							break
						}
						leftBotI = top(leftMoves)
						if data.Position < robot_data_list[leftBotI].Position {
							robot_data_list = leftedHealths(robot_data_list, index, leftBotI)
							if robot_data_list[leftBotI].Health == 0 {
								leftMoves = pop(leftMoves)
							}
						}
					}
					if robot_data_list[index].Health > 0 {
						leftMoves = append(leftMoves, index)
					}
				} else {
					rightMoves = append(rightMoves, index)
				}
			} else if data.Direction == "L" && isNotEmpty(rightMoves) {
				rightBotI = top(rightMoves)
				if data.Position > robot_data_list[rightBotI].Position {
					for true {
						if data.Position < robot_data_list[rightBotI].Position || !isNotEmpty(rightMoves) || robot_data_list[index].Health == 0 {
							break
						}
				        rightBotI = top(rightMoves)
						if data.Position > robot_data_list[rightBotI].Position {
							robot_data_list = leftedHealths(robot_data_list, index, rightBotI)
							if robot_data_list[rightBotI].Health == 0 {
								rightMoves = pop(rightMoves)
							}
						}
					}
					if robot_data_list[index].Health > 0 {
						leftMoves = append(leftMoves, index)
					}
				} else {
					leftMoves = append(leftMoves, index)
				}
			} else if data.Direction == "R" && !isNotEmpty(leftMoves) {
				rightMoves = append(rightMoves, index)
			} else {
				leftMoves = append(leftMoves, index)
			}
		}
		// fmt.Printf("data array: %v\n", robot_data_list)
		// fmt.Printf("left stack: %v ", leftMoves)
		// fmt.Printf("right stack: %v\n", rightMoves)

	}
	// fmt.Printf("after computing: %v\n", robot_data_list)
	// prepare to answer
	sort.SliceStable(robot_data_list, func(k, j int) bool { // sort the robot data list by id
		return robot_data_list[k].id < robot_data_list[j].id
	})
	for j := 0; j < len_healths; j++ {
		if robot_data_list[j].Health > 0 {
			result = append(result, robot_data_list[j].Health)
		}
	}
	return result
}

func main() {
	directions := "RRRRR"
	positions := []int {5,4,3,2,1}
	healths := []int {2,17,9,15,10}
	fmt.Println(survivedRobotsHealths(positions, healths, directions))
}