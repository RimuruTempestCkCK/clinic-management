<script>
    import { onMount } from 'svelte';
    import { fetchWithAuth } from '$lib/api';
    
    let data = $state([]);
    let loading = $state(true);
    let error = $state('');

    onMount(async () => {
        try {
            data = await fetchWithAuth('/medical-records');
        } catch (e) {
            error = e.message;
        } finally {
            loading = false;
        }
    });
</script>

<svelte:head>
    <title>Medical Records · Clinic Management</title>
</svelte:head>

<section class="hero">
    <div class="hero-text">
        <h1 class="hero-title">Medical Records</h1>
        <p class="hero-sub">Manage all medical records data from this page.</p>
    </div>
    <div class="hero-actions">
        <button class="btn btn--primary">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M12 5v14M5 12h14"/></svg> 
            New Record
        </button>
    </div>
</section>

<div class="grid">
    <section class="col-12 card">
        <div class="card-head">
            <div class="card-title-wrap">
                <h2 class="card-title">Medical Records List</h2>
            </div>
        </div>
        
        {#if loading}
            <div style="padding: 24px; color: var(--t-muted);">Loading data...</div>
        {:else if error}
            <div style="padding: 24px; color: var(--danger);">{error}</div>
        {:else}
            <table class="table">
                <thead>
                    <tr>
                        <th>Date</th>
                        <th>Patient</th>
                        <th>Doctor</th>
                        <th>Diagnosis</th>
                        <th>Treatment</th>
                    </tr>
                </thead>
                <tbody>
                    {#each data as item}
                        <tr>
                            <td class="cell-date">{new Date(item.CreatedAt).toLocaleDateString()}</td>
                            <td class="cell-name">{item.Patient?.Name || `ID: ${item.PatientID}`}</td>
                            <td>{item.Doctor?.Name || `ID: ${item.DoctorID}`}</td>
                            <td>{item.Diagnosis}</td>
                            <td>{item.Treatment}</td>
                        </tr>
                    {/each}
                    {#if data.length === 0}
                        <tr><td colspan="5" style="text-align: center; color: var(--t-muted); padding: 1rem;">No medical records found.</td></tr>
                    {/if}
                </tbody>
            </table>
        {/if}
    </section>
</div>
