<script>
    import { login } from '$lib/api';
    import { goto } from '$app/navigation';
    
    let email = 'admin@clinic.com';
    let password = 'password123';
    let error = '';
    
    async function handleLogin() {
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
    <form on:submit|preventDefault={handleLogin} class="space-y-4">
        <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Email</label>
            <input type="email" bind:value={email} class="w-full border-gray-300 rounded-md shadow-sm p-2 border" required />
        </div>
        <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Password</label>
            <input type="password" bind:value={password} class="w-full border-gray-300 rounded-md shadow-sm p-2 border" required />
        </div>
        <button type="submit" class="w-full bg-blue-600 text-white font-semibold p-2 rounded-md hover:bg-blue-700 transition">
            Log In
        </button>
    </form>
</div>
