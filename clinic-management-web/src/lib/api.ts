const API_URL = 'http://localhost:8080/api/v1';

export async function login(email, password) {
    const res = await fetch(`${API_URL}/auth/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password })
    });
    if (!res.ok) throw new Error('Login failed');
    const data = await res.json();
    if (typeof window !== 'undefined') {
        localStorage.setItem('token', data.token);
    }
    return data;
}

export function logout() {
    if (typeof window !== 'undefined') {
        localStorage.removeItem('token');
    }
}

export async function fetchWithAuth(endpoint, options = {}) {
    const token = typeof window !== 'undefined' ? localStorage.getItem('token') : null;
    const headers = {
        'Content-Type': 'application/json',
        ...options.headers,
    };
    if (token) {
        headers['Authorization'] = `Bearer ${token}`;
    }
    
    const res = await fetch(`${API_URL}${endpoint}`, {
        ...options,
        headers
    });
    
    if (res.status === 401) {
        logout();
        if (typeof window !== 'undefined') window.location.href = '/login';
        throw new Error('Unauthorized');
    }
    
    return res.json();
}
