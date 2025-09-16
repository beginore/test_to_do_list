import { useState, useEffect } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { Add, List, Toggle, Delete, DeleteAll } from "../wailsjs/go/main/App";

function App() {
    const [todos, setTodos] = useState([]);
    const [todoInput, setTodoInput] = useState('');
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');
    const [filter, setFilter] = useState('all'); // all | active | completed
    const [theme, setTheme] = useState('light');

    useEffect(() => {
        loadTodos();
    }, []);

    async function loadTodos() {
        setLoading(true);
        setError('');
        try {
            const data = await List();
            setTodos(Array.isArray(data) ? data : []);
        } catch (err) {
            setError('Failed to load todos: ' + err.message);
            setTodos([]);
        } finally {
            setLoading(false);
        }
    }

    async function handleAdd() {
        if (todoInput.trim() === '') return;
        try {
            await Add(todoInput);
            setTodoInput('');
            await loadTodos();
        } catch (err) {
            setError('Failed to add todo: ' + err.message);
        }
    }

    async function handleToggle(id) {
        try {
            await Toggle(id);
            await loadTodos();
        } catch (err) {
            setError('Failed to update todo: ' + err.message);
        }
    }

    async function handleDelete(id) {
        if (!window.confirm("Are you sure you want to delete this todo?")) return;
        try {
            await Delete(id);
            await loadTodos();
        } catch (err) {
            setError('Failed to delete todo: ' + err.message);
        }
    }

    async function handleDeleteAll() {
        if (!window.confirm("Are you sure you want to delete ALL todos?")) return;
        try {
            await DeleteAll();
            await loadTodos();
        } catch (err) {
            setError('Failed to clear todos: ' + err.message);
        }
    }

    function formatDate(dateString) {
        if (!dateString) return '';
        try {
            const date = new Date(dateString);
            return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
        } catch {
            return dateString;
        }
    }

    function toggleTheme() {
        setTheme(prev => prev === 'light' ? 'dark' : 'light');
    }

    // Фильтрация задач
    const filteredTodos = todos.filter(todo => {
        if (filter === 'active') return !todo.Completed;
        if (filter === 'completed') return todo.Completed;
        return true;
    });

    return (
        <div id="App" className={`app-container ${theme}`}>
            <img src={logo} id="logo" alt="logo" />
            <div className="header">
                <h1>Todo List</h1>
                <button onClick={toggleTheme}>
                    {theme === 'light' ? '🌙 Dark' : '☀️ Light'}
                </button>
            </div>
            {/* Панель управления */}
            <div className="todo-input-container">
                <input
                    type="text"
                    value={todoInput}
                    onChange={(e) => setTodoInput(e.target.value)}
                    placeholder="Enter todo"
                    onKeyPress={(e) => e.key === 'Enter' && handleAdd()}
                    disabled={loading}
                />
                <button onClick={handleAdd} disabled={loading}>
                    {loading ? 'Adding...' : 'Add'}
                </button>
                <button onClick={handleDeleteAll} disabled={loading || todos.length === 0}>
                    Clear All
                </button>
                <button onClick={loadTodos} disabled={loading}>
                    Refresh
                </button>
            </div>

            {/* Фильтры */}
            <div className="filter-container">
                <button onClick={() => setFilter('all')} className={filter === 'all' ? 'active' : ''}>All</button>
                <button onClick={() => setFilter('active')} className={filter === 'active' ? 'active' : ''}>Active</button>
                <button onClick={() => setFilter('completed')} className={filter === 'completed' ? 'active' : ''}>Completed</button>
            </div>

            {/* Ошибки */}
            {error && <div className="error-message">{error}</div>}

            {/* Список */}
            {loading ? (
                <div className="loading">Loading todos...</div>
            ) : (
                <ul className="todo-list">
                    {filteredTodos.map((todo) => (
                        <li key={todo.ID} className="todo-item">
                            <div className="todo-content">
                                <span
                                    onClick={() => handleToggle(todo.ID)}
                                    className={`todo-text ${todo.Completed ? 'completed' : ''}`}
                                >
                                    {todo.Title}
                                </span>
                                <div className="todo-dates">
                                    <small>Created: {formatDate(todo.CreatedAt)}</small>
                                    {todo.Completed && todo.CompletedAt && (
                                        <small>Completed: {formatDate(todo.CompletedAt)}</small>
                                    )}
                                </div>
                            </div>
                            <button
                                onClick={() => handleDelete(todo.ID)}
                                className="delete-btn"
                                disabled={loading}
                                title="Delete todo"
                            >
                                ❌
                            </button>
                        </li>
                    ))}
                </ul>
            )}

            {!loading && filteredTodos.length === 0 && !error && (
                <div className="empty-state">No todos here 🚀</div>
            )}
        </div>
    );
}

export default App;
