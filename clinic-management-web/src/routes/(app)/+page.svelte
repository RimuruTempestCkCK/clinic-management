<script>
    import { fetchWithAuth } from '$lib/api';
    import { onMount } from 'svelte';
    
    let doctors = $state([]);
    let appointments = $state([]);
    let loading = $state(true);
    let today = new Date().toLocaleDateString('en-US', { weekday: 'long', month: 'long', day: 'numeric', year: 'numeric' });
    
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

<svelte:head>
    <title>Dashboard · Clinic Management</title>
</svelte:head>

<section class="hero">
    <div class="hero-text">
        <span class="eyebrow" id="heroDate">{today}</span>
        <h1 class="hero-title">Welcome back, <span class="accent">Admin</span></h1>
        <p class="hero-sub">Here is the overview of your clinic's operations. The data below is fetched in real-time from the API.</p>
    </div>
    <div class="hero-actions">
        <button class="btn btn--primary">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M12 5v14M5 12h14"/></svg> 
            New Appointment
        </button>
    </div>
</section>

{#if loading}
    <div style="padding: 2rem; text-align: center; color: var(--t-muted);">Loading data...</div>
{:else}
    <section class="kpi-grid" aria-label="Key metrics">
        <article class="kpi-card c-success">
            <div class="kpi-top">
                <div class="kpi-identity">
                    <div class="kpi-icon success">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><path d="M14 2v6h6M16 13H8M16 17H8M10 9H8"/></svg>
                    </div>
                    <div class="kpi-label">Appointments</div>
                </div>
            </div>
            <div class="kpi-value">{appointments.length}</div>
            <div class="kpi-compare">Total appointments in the system</div>
        </article>

        <article class="kpi-card c-primary">
            <div class="kpi-top">
                <div class="kpi-identity">
                    <div class="kpi-icon primary">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M8 12h8M12 8v8"/></svg>
                    </div>
                    <div class="kpi-label">Doctors</div>
                </div>
            </div>
            <div class="kpi-value">{doctors.length}</div>
            <div class="kpi-compare">Registered medical staff</div>
        </article>
    </section>

    <div class="grid">
        <section class="col-6 card">
            <div class="card-head">
                <div class="card-title-wrap">
                    <span class="eyebrow">Schedule</span>
                    <h2 class="card-title">Recent Appointments</h2>
                </div>
                <a class="card-action" href="#">View all <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M5 12h14M13 5l7 7-7 7"/></svg></a>
            </div>
            <table class="table">
                <thead>
                    <tr>
                        <th>Patient ID</th>
                        <th>Status</th>
                        <th>Date</th>
                    </tr>
                </thead>
                <tbody>
                    {#each appointments.slice(0, 5) as app}
                        <tr>
                            <td class="cell-name">#{app.patient_id}</td>
                            <td>
                                {#if app.status === 'CONFIRMED'}
                                    <span class="tag t-new">{app.status}</span>
                                {:else}
                                    <span class="tag t-used">{app.status}</span>
                                {/if}
                            </td>
                            <td class="cell-date">{new Date(app.schedule_date).toLocaleDateString()}</td>
                        </tr>
                    {/each}
                    {#if appointments.length === 0}
                        <tr><td colspan="3" style="text-align: center; color: var(--t-muted);">No recent appointments</td></tr>
                    {/if}
                </tbody>
            </table>
        </section>

        <section class="col-6 card">
            <div class="card-head">
                <div class="card-title-wrap">
                    <span class="eyebrow">Staff</span>
                    <h2 class="card-title">Doctors Roster</h2>
                </div>
                <a class="card-action" href="#">Manage <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M5 12h14M13 5l7 7-7 7"/></svg></a>
            </div>
            <ul class="todo-list">
                {#each doctors.slice(0, 5) as doc}
                    <li class="todo-item" style="display: flex; justify-content: space-between; align-items: center; padding: 12px 16px;">
                        <div>
                            <div class="todo-text" style="font-weight: 500;">{doc.name}</div>
                            <div style="font-size: 12px; color: var(--t-muted);">{doc.specialization?.name || 'General Practitioner'}</div>
                        </div>
                        <span class="todo-badge upcoming">{doc.phone}</span>
                    </li>
                {/each}
                {#if doctors.length === 0}
                    <li class="todo-item" style="padding: 12px 16px; color: var(--t-muted);">No doctors available</li>
                {/if}
            </ul>
        </section>
    </div>
{/if}
