<script>
    import { onMount } from 'svelte';
    import { fetchWithAuth } from '$lib/api';
    
    let data = $state([]);
    let loading = $state(true);
    let error = $state('');

    onMount(async () => {
        try {
            data = await fetchWithAuth('/medicines');
        } catch (e) {
            error = e.message;
        } finally {
            loading = false;
        }
    });
</script>

<svelte:head>
    <title>Medicines · Clinic Management</title>
</svelte:head>

<section class="hero">
    <div class="hero-text">
        <h1 class="hero-title">Medicines</h1>
        <p class="hero-sub">Manage all medicines data from this page.</p>
    </div>
    <div class="hero-actions">
        <button class="btn btn--primary">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M12 5v14M5 12h14"/></svg> 
            New Medicine
        </button>
    </div>
</section>

<div class="grid">
    <section class="col-12 card">
        <div class="card-head">
            <div class="card-title-wrap">
                <h2 class="card-title">Medicines Inventory</h2>
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
                        <th>Description</th>
                        <th style="text-align: right">Price</th>
                        <th>Stock</th>
                    </tr>
                </thead>
                <tbody>
                    {#each data as item}
                        <tr>
                            <td class="cell-name">{item.Name}</td>
                            <td>{item.Description}</td>
                            <td class="cell-price pos" style="text-align: right">${item.Price.toFixed(2)}</td>
                            <td>
                                {#if item.Stock > 10}
                                    <span class="tag t-new">{item.Stock} units</span>
                                {:else if item.Stock > 0}
                                    <span class="tag t-used">{item.Stock} units (Low)</span>
                                {:else}
                                    <span class="tag t-unavail">Out of Stock</span>
                                {/if}
                            </td>
                        </tr>
                    {/each}
                    {#if data.length === 0}
                        <tr><td colspan="4" style="text-align: center; color: var(--t-muted); padding: 1rem;">No medicines found.</td></tr>
                    {/if}
                </tbody>
            </table>
        {/if}
    </section>
</div>
