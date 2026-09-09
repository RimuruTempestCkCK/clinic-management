<script>
    import { login } from '$lib/api';
    import { goto } from '$app/navigation';
    
    let email = $state('admin@clinic.com');
    let password = $state('password123');
    let error = $state('');
    let loading = $state(false);
    
    async function handleLogin(event) {
        event.preventDefault();
        loading = true;
        try {
            await login(email, password);
            goto('/');
        } catch (e) {
            error = e.message;
        } finally {
            loading = false;
        }
    }
</script>

<svelte:head>
    <title>Sign in · Clinic Management</title>
</svelte:head>

<div class="auth-shell" style="justify-content: center; background-color: var(--c-bg);">
    <main class="auth-main" style="flex: none; max-width: 500px; width: 100%; border-radius: 12px; box-shadow: 0 4px 6px rgba(0,0,0,0.1);">
        <div class="auth-main-top">
            <a href="/" style="font-size:12.5px;color:var(--t-muted);display:inline-flex;align-items:center;gap:6px">
                <svg viewBox="0 0 24 24" width="13" height="13" fill="none" stroke="currentColor" stroke-width="2"><path d="M19 12H5M12 19l-7-7 7-7"/></svg> Back to home
            </a>
            <div class="switch-link">New here? <a href="#">Create account</a></div>
        </div>
        
        <div class="auth-card">
            <h2>Welcome back</h2>
            <p class="sub">Sign in to your Clinic workspace to pick up where you left off.</p>
            
            {#if error}
                <div style="background-color: var(--danger-light); color: var(--danger); padding: 12px; border-radius: 6px; margin-bottom: 16px; font-size: 14px;">
                    {error}
                </div>
            {/if}

            <form class="auth-form" onsubmit={handleLogin}>
                <div class="field">
                    <label class="field-label" for="email">Email</label>
                    <div class="input-icon">
                        <span class="ico">
                            <svg viewBox="0 0 24 24"><rect x="3" y="5" width="18" height="14" rx="2"/><path d="m3 7 9 6 9-6"/></svg>
                        </span>
                        <input id="email" class="input" type="email" bind:value={email} placeholder="you@company.com" autocomplete="email" required>
                    </div>
                </div>
                
                <div class="field">
                    <div class="field-row">
                        <label class="field-label" for="password">Password</label> 
                        <a href="#">Forgot?</a>
                    </div>
                    <div class="input-icon">
                        <span class="ico">
                            <svg viewBox="0 0 24 24"><rect x="3" y="11" width="18" height="11" rx="2"/><path d="M7 11V7a5 5 0 0 1 10 0v4"/></svg>
                        </span>
                        <input id="password" class="input" type="password" bind:value={password} placeholder="••••••••" autocomplete="current-password" required>
                    </div>
                </div>
                
                <label class="check">
                    <input type="checkbox" checked="checked"> <span class="box"></span> Keep me signed in for 30 days
                </label> 
                
                <button class="btn btn--primary auth-submit" type="submit" disabled={loading}>
                    {loading ? 'Signing in...' : 'Sign in'} 
                    <svg viewBox="0 0 24 24"><path d="M5 12h14M13 5l7 7-7 7"/></svg>
                </button>
            </form>
        </div>
        <div class="auth-main-bottom">By signing in you agree to our <a href="#">Terms</a> and <a href="#">Privacy Policy</a>.</div>
    </main>
</div>
