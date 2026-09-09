<script>
    import { onMount } from 'svelte';
    import { fetchWithAuth } from '$lib/api';
    
    let data = $state([]);
    let loading = $state(true);
    let error = $state('');

    onMount(async () => {
        try {
            data = await fetchWithAuth('/doctors');
        } catch (e) {
            error = e.message;
        } finally {
            loading = false;
        }
    });
</script>

<svelte:head>
    <title>Doctors · Clinic Management</title>
</svelte:head>

<section class="hero">
    <div class="hero-text">
        <h1 class="hero-title">Doctors</h1>
        <p class="hero-sub">Manage all doctors data from this page.</p>
    </div>
    <div class="hero-actions">
        <button class="btn btn--primary">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M12 5v14M5 12h14"/></svg> 
            New Doctor
        </button>
    </div>
</section>

<div class="grid">
    <section class="col-12 card">
        <div class="card-head">
            <div class="card-title-wrap">
                <h2 class="card-title">Doctors List</h2>
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
                        <th>Name</th>
                        <th>Specialization</th>
                        <th>Phone</th>
                        <th>Bio</th>
                    </tr>
                </thead>
                <tbody>
                    {#each data as item}
                        <tr>
                            <td class="cell-name">{item.Name}</td>
                            <td><span class="tag t-used">{item.Specialization?.Name || 'General'}</span></td>
                            <td>{item.Phone}</td>
                            <td>{item.Bio}</td>
                        </tr>
                    {/each}
                    {#if data.length === 0}
                        <tr><td colspan="4" style="text-align: center; color: var(--t-muted); padding: 1rem;">No doctors found.</td></tr>
                    {/if}
                </tbody>
            </table>
        {/if}
    </section>
</div>
