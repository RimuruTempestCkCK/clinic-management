<script>
    import { onMount } from 'svelte';
    import { fetchWithAuth } from '$lib/api';
    
    let data = $state([]);
    let loading = $state(true);
    let error = $state('');

    onMount(async () => {
        try {
            data = await fetchWithAuth('/appointments');
        } catch (e) {
            error = e.message;
        } finally {
            loading = false;
        }
    });
</script>

<svelte:head>
    <title>Appointments · Clinic Management</title>
</svelte:head>

<section class="hero">
    <div class="hero-text">
        <h1 class="hero-title">Appointments</h1>
        <p class="hero-sub">Manage all appointments data from this page.</p>
    </div>
    <div class="hero-actions">
        <!-- <button class="btn btn--primary">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M12 5v14M5 12h14"/></svg> 
            New Appointment
        </button> -->
    </div>
</section>

<div class="grid">
    <section class="col-12 card">
        <div class="card-head">
            <div class="card-title-wrap">
                <h2 class="card-title">Appointments List</h2>
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
                        <th>Patient Name</th>
                        <th>Doctor</th>
                        <th>Date & Time</th>
                        <th>Status</th>
                    </tr>
                </thead>
                <tbody>
                    {#each data as item}
                        <tr>
                            <td class="cell-name">{item.Patient?.Name || `ID: ${item.PatientID}`}</td>
                            <td>{item.Doctor?.Name || `ID: ${item.DoctorID}`}</td>
                            <td class="cell-date">{new Date(item.ScheduleDate).toLocaleString()}</td>
                            <td>
                                <span class="tag {item.Status === 'CONFIRMED' ? 't-new' : item.Status === 'COMPLETED' ? 't-used' : 't-unavail'}">
                                    {item.Status}
                                </span>
                            </td>
                        </tr>
                    {/each}
                    {#if data.length === 0}
                        <tr><td colspan="4" style="text-align: center; color: var(--t-muted); padding: 1rem;">No appointments found.</td></tr>
                    {/if}
                </tbody>
            </table>
        {/if}
    </section>
</div>
