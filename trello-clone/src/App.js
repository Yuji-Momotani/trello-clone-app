import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import Auth from './components/Auth/Auth';
import { useEffect } from 'react';
import axios from 'axios';
import TrelloApp from "./components/trello-app/TodoApp";

const App = () => {
	useEffect(() => {
		const getCsrf = async() => {
			axios.defaults.withCredentials = true;
			const {data} = await axios.get(
				`${process.env.REACT_APP_API_URL}/csrf`,
				// {withCredentials: true}
			)
			axios.defaults.headers.common["X-CSRF-Token"] = data.csrf_token;
		}
		getCsrf()
	})
	return (
		<Router>
			<Routes>
				<Route path="/" element={<Auth/ >} />
				{/* 以下の参考例をもとに実際の機能のルーティングを行う */}
				{/* <Route path="/todo" element={<TodoApp/>} /> */}

				<Route path="/task" element={<TrelloApp />} />
			</Routes>
		</Router>
	);
}

export default App;