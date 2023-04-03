package ws

type AcceptActionType string

/*
 * Action types are the types of actions that can be performed by the server
 * These are the actions that the client can send to the server
 * The server can also send actions to the client
 *	***  Naming Convention Followed	***
 		1. Must start with the word ACTION_
		2. Must be all in Block Letters
		3. Must be separated by underscores
		4. Full type format be in ACTION_<WORK_TO_BE_DONE>_<OPTIONAL_SUB_WORK>
 * Ex: ACTION_CREATE_OPD_APPOINTMENT, ACTION_LEAVE_USER_ROOM
*/

const (
	ACTION_LEAVE_USER_ROOM AcceptActionType = "ACTION_LEAVE_USER_ROOM"
)
