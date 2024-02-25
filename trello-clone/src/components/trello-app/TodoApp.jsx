import './TodoApp.css';
import { Header } from "../header/Header";
import { TaskCards } from "../task/TaskCards";

const TodoApp = () => {
	return (
		<div className="app">
			<Header />
			<TaskCards />
		</div>
	)
}

export default TodoApp;
