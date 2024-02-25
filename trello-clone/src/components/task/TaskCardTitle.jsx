import React, { useContext, useEffect, useState } from 'react'
import { TaskCardContext } from "./TaskCard";
import { CardAndTaksContext } from "./TaskCards";
import taskCardUtil from "../../util/taskCard";

export const TaskCardTitle = ({title}) => {
	const [isClick, setIsClick] = useState(false);
	const [isUpdate, setIsUpdate] = useState(false);
	const [inputCardTitle, setInputCardTitle] = useState(title);
	const {updateCard} = taskCardUtil();
	const taskCard = useContext(TaskCardContext);
	const [,setTaskCardsList] = useContext(CardAndTaksContext);

	useEffect(() => {
		const doUpdateCard = async() => {
			if (!isUpdate) {
				return;
			}
			try {
				await updateCard(taskCard.id, inputCardTitle, taskCard.sort_no);
				setTaskCardsList(prevCards => {
					const newCards = prevCards.map(prevCard => {
						if (prevCard.id === taskCard.id) {
							prevCard.title = inputCardTitle;
						}
						return prevCard;
					});
					return newCards;
				});
			} catch(e) {
				// 一旦放置
				console.log(e);
			} finally {
				setIsUpdate(false);
			}
		}
		doUpdateCard();
	}, [isClick])

	const handleClick = () => {
		setIsClick(true);
	};

	const handleChange = async(e) => {
		setInputCardTitle(e.target.value); 
	};

	const handleSubmit = (e) => {
		e.preventDefault();
		setIsClick(false);
		setIsUpdate(true);
	}

	const textBoxOnBlur = () => {
		setIsClick(false);
		setIsUpdate(true);
	}

	return (
		<div onClick={handleClick} className="taskCardTitleInputArea">
			{
				isClick ? (
					<form onSubmit={(e) => handleSubmit(e)}>
						<input autoFocus
							className="taskCardTitleInput"
							type="text" value={inputCardTitle} 
							onChange={(e) => handleChange(e)} 
							onBlur={textBoxOnBlur}
							maxLength="10"
						/>
					</form>
				) 
				:
				<h3>{inputCardTitle}</h3>
			}
		</div>
	)
}
