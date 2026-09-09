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

<div class="auth-shell">
    <aside class="auth-aside" style="background-image: url('/clinic-bg.jpg'); background-size: cover; background-position: center; position: relative;">
        <!-- Add a subtle dark overlay so text is readable if we want, or just leave it blank -->
        <div style="position: absolute; inset: 0; background: linear-gradient(to bottom, rgba(0,0,0,0.3) 0%, rgba(0,0,0,0.7) 100%);"></div>
        
        <div class="auth-brand" style="position: relative; z-index: 1;">
            <div class="logo">
                <svg viewBox="0 0 36 36" xmlns="http://www.w3.org/2000/svg"><path fill="#fff" d="M14.747 9.125c.527-1.426 1.736-2.573 3.317-2.573c1.643 0 2.792 1.085 3.318 2.573l6.077 16.867c.186.496.248.931.248 1.147c0 1.209-.992 2.046-2.139 2.046c-1.303 0-1.954-.682-2.264-1.611l-.931-2.915h-8.62l-.93 2.884c-.31.961-.961 1.642-2.232 1.642c-1.24 0-2.294-.93-2.294-2.17c0-.496.155-.868.217-1.023l6.233-16.867zm.34 11.256h5.891l-2.883-8.992h-.062l-2.946 8.992z"/></svg>
            </div>
            <div class="name">Clinic Web</div>
        </div>
        
        <div class="auth-aside-body" style="position: relative; z-index: 1;">
            <span class="auth-aside-eyebrow" style="color: rgba(255,255,255,0.8);">Clinic Management System</span>
            <h1 style="color: white; text-shadow: 0 2px 4px rgba(0,0,0,0.5);">Modern solutions for modern healthcare.</h1>
            <p style="color: rgba(255,255,255,0.9); text-shadow: 0 1px 2px rgba(0,0,0,0.5);">Faster medical records, cleaner appointments, and a design system that scales for your operational needs.</p>
        </div>
        
        <div class="auth-aside-footer" style="position: relative; z-index: 1; color: rgba(255,255,255,0.6);">
            <span>© 2026</span> <span>BUILT IN CLINIC</span>
        </div>
    </aside>

    <main class="auth-main">
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
