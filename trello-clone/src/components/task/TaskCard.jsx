import React, { createContext, useEffect, useState } from 'react'
import { TaskCardTitle } from "./TaskCardTitle"
import { TaskCardDeleteButton } from "./button/TaskCardDeleteButton"
import { TaskAddInput } from "./input/TaskAddInput"
import { Tasks } from "./Tasks"
import { Draggable } from "react-beautiful-dnd"

export const TaskListContext = createContext();
export const TaskCardContext = createContext();

export const TaskCard = ({taskCard, draggableIndex}) => {
	const [inputText, setInputText] = useState("");
	const [taskList, setTaskList] = useState(taskCard.tasks);


	return (
		<Draggable draggableId={`taskCard-${taskCard.id}`} index={draggableIndex}>
			{(provided) => (
				<TaskCardContext.Provider value={taskCard}>
					<TaskListContext.Provider value={[taskList, setTaskList]}>
						<div className="taskCard" ref={provided.innerRef} {...provided.draggableProps}>
							<div className="taskCardTitleAndTaskCardDeleteButtonArea" {...provided.dragHandleProps}>
								<TaskCardTitle title={taskCard.title} />
								<TaskCardDeleteButton 
									taskCard={taskCard} 
									// setTaskCardsList={setTaskCardsList}
								/>
							</div>
							<TaskAddInput 
								inputText={inputText} 
								setInputText={setInputText} 
							/>
							<Tasks />
						</div>
					</TaskListContext.Provider>
				</TaskCardContext.Provider>
			)}
		</Draggable>
	)
}
