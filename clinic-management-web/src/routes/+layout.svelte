<script>
    import '../app.css';
</script>

<div class="min-h-screen bg-gray-50 text-gray-900 font-sans">
    <nav class="bg-white shadow-sm p-4 flex justify-between items-center">
        <h1 class="text-xl font-bold text-blue-600">🏥 Clinic Manager</h1>
        <div class="space-x-4">
            <a href="/" class="text-gray-600 hover:text-blue-500">Dashboard</a>
            <a href="/login" class="text-gray-600 hover:text-blue-500">Login</a>
            <button on:click={() => { localStorage.removeItem('token'); window.location.href='/login' }} class="text-red-500">Logout</button>
        </div>
    </nav>
    <main class="p-8 max-w-5xl mx-auto">
        <slot />
    </main>
</div>
