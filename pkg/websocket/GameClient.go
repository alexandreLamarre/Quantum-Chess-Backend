package websocket

import (
	"fmt"
	"github.com/alexandreLamarre/Quantum-Chess-Backend/pkg/quantumchess"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
)
//DEBUG_DECODE toggles decoding debug messages in the console.
var DEBUG_DECODE = true

// GameClient represents a new websocket connection to manage games, with the user's id, and the game pool it is connected to.
type GameClient struct {
	ID       string
	Conn     *websocket.Conn
	GamePool *GamePool
}

//GameMessage allows us to unpack the content of JSON transmitted through the game pool.
type GameMessage struct {
	Type          int                        `json:"type"` // 0 = player connected, 1 = board update, 2 = message, 3= opponent leave, 4= spectator join/leave
	Pid           string                     `json:"pid"`
	Color         int                        `json:"color"`
	GameStart     bool                       `json:"start"`
	GameEnd       bool                       `json:"end"`
	Message       string                     `json:"message"`
	Move          [2]int                     `json:"move"`
	Board         [64]int   	  			  `json:"board"`
	Entanglements map[string]interface{} 	  `json:"entanglements"`
	Pieces        map[string]interface{}      `json:"pieces"`
	NewBoard 	  [64]int 					  `json:"newBoard"`
	NewPieces     quantumchess.Pieces         `json:"newPieces"`
	NewEntanglements quantumchess.Entanglements `json:"newEntanglements"`
}

//GameRead performs the parsing of JSON message, and then appropriately broadcasts the transformed message to other users.
func (c *GameClient) GameRead() {
	defer func() {
		c.GamePool.Unregister <- c
		c.Conn.Close()
	}()

	for {

		message := &GameMessage{}
		err := c.Conn.ReadJSON(message)
		if err != nil {
			fmt.Println("\n message: \n", message)
			log.Println(err)
			return
		}
		if message.Type == 2{
			//c.GamePool.Broadcast <- *message
		} else if message.Type == 1{
			board := &quantumchess.Board{}
			pieces := &quantumchess.Pieces{}
			entanglements := &quantumchess.Entanglements{}

			decodeBoardFromMessage(message, board, pieces, entanglements)
			if DEBUG_DECODE{ fmt.Println("Moving piece from ", message.Move[0], " to ", message.Move[1])}

			err := quantumchess.ApplyMove(board, entanglements, pieces, message.Move[0], message.Move[1])
			if err != nil{
				fmt.Println("Error applying move")
				log.Println(err)
			}
			var newBoard [64]int
			copy(newBoard[:], board.Positions)
			updateMessage := GameMessage{
				Type:1,
				NewBoard: newBoard,
				NewPieces: *pieces,
				NewEntanglements: *entanglements,

			}
			c.GamePool.Broadcast <- updateMessage
		}


	}
}


func decodeBoardFromMessage(message *GameMessage, board *quantumchess.Board,
	pieces *quantumchess.Pieces, entanglements *quantumchess.Entanglements ){
	board.Positions = message.Board[:]
	pieces.List = make(map[int]*quantumchess.Piece)
	entanglements.List = make(map[int]*quantumchess.Entanglement)

	for id,interf := range message.Pieces{
		newPiece := &quantumchess.Piece{}
		data:= interf.(map[string]interface{})
		for k, v := range data{
			if k == "action"{
				newPiece.Action = v.(string)
			} else if k == "color"{
				newPiece.Color = int(v.(float64))
			} else if k == "moved"{
				newPiece.Moved = v.(bool)
			} else if k == "initialState"{
				initialStateParsed := make(map[string][2]float64)
				initialStateData := v.(map[string]interface{})
				for state, cmplx := range initialStateData{
					var cmplxParsed [2]float64
					cmplxData := cmplx.([]interface{})
					for i, num := range cmplxData{
						cmplxParsed[i] = num.(float64)
					}
					initialStateParsed[state] = cmplxParsed
				}
				newPiece.State = initialStateParsed

			} else if k == "stateSpace"{
				stateSpaceData := v.([]interface{})
				var stateSpaceParse = make([]string, 0,0)
				for _, v := range stateSpaceData{
					stateSpaceParse = append(stateSpaceParse, v.(string))
				}
				newPiece.StateSpace = stateSpaceParse
				//fmt.Println("stateSpaceData",stateSpaceData)
			} else if k == "states"{
				statesParsed := make(map[string][2]float64)
				statesData := v.(map[string]interface{})
				for state, cmplx := range statesData{
					var cmplxParsed [2]float64
					cmplxData := cmplx.([]interface{})
					for i, num := range cmplxData{
						cmplxParsed[i] = num.(float64)
					}
					statesParsed[state] = cmplxParsed
				}
				newPiece.InitialState = statesParsed
			}
		}

		pid,err := strconv.ParseInt(id, 10, 0)
		if err != nil{
			log.Println(err)
		}
		pieces.List[int(pid)] = newPiece
	}


	for id, interf := range message.Entanglements{
		newEntanglement := &quantumchess.Entanglement{}

		if interf == nil {
			newEntanglement = nil
		} else {
			data := interf.(map[string]interface{})
			for attr, value := range data{
				if attr == "elements"{
					entangledId := value.([]int)
					newEntanglement.Elements = entangledId
				} else if attr == "state"{
					entangledState := value.([][2]float64)
					// might have parse it as interface{}
					newEntanglement.State = entangledState
				}
			}
		}


		pid, err := strconv.ParseInt(id, 10, 0)
		if err != nil{
			log.Println(err)
		}
		entanglements.List[int(pid)] = newEntanglement
	}

	if DEBUG_DECODE {
		for k,v:= range pieces.List{
			fmt.Println(k,v)
		}

		for k,v := range entanglements.List{
			fmt.Println(k,v)
		}
	}
}
