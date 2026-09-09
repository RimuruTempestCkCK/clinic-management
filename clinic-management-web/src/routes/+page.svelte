<script>
    import { fetchWithAuth } from '$lib/api';
    import { onMount } from 'svelte';
    
    let doctors = $state([]);
    let appointments = $state([]);
    let loading = $state(true);
    
    onMount(async () => {
        if (!localStorage.getItem('token')) {
            window.location.href = '/login';
            return;
        }
        try {
            doctors = await fetchWithAuth('/doctors');
            appointments = await fetchWithAuth('/appointments');
        } catch (e) {
            console.error(e);
        } finally {
            loading = false;
        }
    });
</script>

<h1 class="text-3xl font-bold mb-8">Dashboard Overview</h1>

{#if loading}
    <div class="animate-pulse flex space-x-4">
        <div class="h-10 bg-gray-200 rounded w-1/4"></div>
        <div class="h-10 bg-gray-200 rounded w-1/4"></div>
    </div>
{:else}
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <!-- Doctors -->
        <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
            <h2 class="text-xl font-semibold mb-4 flex items-center">👨‍⚕️ Doctors ({doctors.length})</h2>
            <ul class="space-y-3">
                {#each doctors.slice(0, 5) as doc}
                    <li class="flex justify-between items-center p-3 hover:bg-gray-50 rounded-lg transition">
                        <div>
                            <p class="font-medium text-gray-900">{doc.name}</p>
                            <p class="text-sm text-gray-500">{doc.specialization?.name || 'General'}</p>
                        </div>
                        <span class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded-full">{doc.phone}</span>
                    </li>
                {/each}
            </ul>
        </div>
        
        <!-- Appointments -->
        <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
            <h2 class="text-xl font-semibold mb-4 flex items-center">📅 Appointments ({appointments.length})</h2>
            <ul class="space-y-3">
                {#each appointments.slice(0, 5) as app}
                    <li class="p-3 hover:bg-gray-50 rounded-lg transition border-l-4 {app.status === 'CONFIRMED' ? 'border-green-500' : 'border-yellow-400'}">
                        <div class="flex justify-between items-start">
                            <div>
                                <p class="font-medium text-gray-900">Patient #{app.patient_id}</p>
                                <p class="text-sm text-gray-500">{new Date(app.schedule_date).toLocaleDateString()}</p>
                            </div>
                            <span class="text-xs font-semibold px-2 py-1 rounded-full {app.status === 'CONFIRMED' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'}">
                                {app.status}
                            </span>
                        </div>
                    </li>
                {/each}
            </ul>
        </div>
    </div>
{/if}
