<script>
    import { login } from '$lib/api';
    import { goto } from '$app/navigation';
    
    let email = $state('admin@clinic.com');
    let password = $state('password123');
    let error = $state('');
    
    async function handleLogin(event) {
        event.preventDefault();
        try {
            await login(email, password);
            goto('/');
        } catch (e) {
            error = e.message;
        }
    }
</script>

<div class="max-w-md mx-auto mt-20 bg-white p-8 rounded-xl shadow-md border border-gray-100">
    <h2 class="text-2xl font-bold mb-6 text-center text-gray-800">Welcome Back</h2>
    {#if error}
        <div class="bg-red-50 text-red-600 p-3 rounded mb-4 text-sm">{error}</div>
    {/if}
    <form onsubmit={handleLogin} class="space-y-4">
        <div>
            <label for="email" class="block text-sm font-medium text-gray-700 mb-1">Email</label>
            <input id="email" type="email" bind:value={email} class="w-full border-gray-300 rounded-md shadow-sm p-2 border" required />
        </div>
        <div>
            <label for="password" class="block text-sm font-medium text-gray-700 mb-1">Password</label>
            <input id="password" type="password" bind:value={password} class="w-full border-gray-300 rounded-md shadow-sm p-2 border" required />
        </div>
        <button type="submit" class="w-full bg-blue-600 text-white font-semibold p-2 rounded-md hover:bg-blue-700 transition">
            Log In
        </button>
    </form>
</div>
