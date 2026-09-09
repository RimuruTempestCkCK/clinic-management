<script>
    import { page } from '$app/stores';

    const navigation = [
        {
            label: "Workspace",
            items: [
                { key: "dashboard", text: "Dashboard", href: "/", icon: '<path d="M3 12 12 3l9 9"/><path d="M5 10v10h14V10"/>' },
                { key: "appointments", text: "Appointments", href: "/appointments", icon: '<rect x="3" y="4" width="18" height="18" rx="2"/><path d="M16 2v4M8 2v4M3 10h18"/>' }
            ]
        },
        {
            label: "Management",
            items: [
                { key: "doctors", text: "Doctors", href: "/doctors", icon: '<circle cx="12" cy="12" r="9"/><path d="M8 12h8M12 8v8"/>' },
                { key: "patients", text: "Patients", href: "/patients", icon: '<path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><path d="M14 2v6h6"/>' },
                { key: "records", text: "Medical Records", href: "/records", icon: '<path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><path d="M14 2v6h6M16 13H8M16 17H8M10 9H8"/>' },
                { key: "medicines", text: "Medicines", href: "/medicines", icon: '<rect x="3" y="8" width="18" height="8" rx="4"/>' }
            ]
        }
    ];

    let currentPath = $derived($page.url.pathname);
</script>

<aside class="d-sidebar">
    <div class="brand">
        <div class="brand-logo">
            <svg viewBox="0 0 36 36" xmlns="http://www.w3.org/2000/svg">
                <path fill="#ffffff" d="M14.747 9.125c.527-1.426 1.736-2.573 3.317-2.573c1.643 0 2.792 1.085 3.318 2.573l6.077 16.867c.186.496.248.931.248 1.147c0 1.209-.992 2.046-2.139 2.046c-1.303 0-1.954-.682-2.264-1.611l-.931-2.915h-8.62l-.93 2.884c-.31.961-.961 1.642-2.232 1.642c-1.24 0-2.294-.93-2.294-2.17c0-.496.155-.868.217-1.023l6.233-16.867zm.34 11.256h5.891l-2.883-8.992h-.062l-2.946 8.992z"/>
            </svg>
        </div>
        <div class="brand-text">
            <div class="brand-name">Clinic Web</div>
            <div class="brand-tag">v1.0.0</div>
        </div>
    </div>
    
    {#each navigation as group}
        <nav class="nav-section">
            <div class="nav-label">{group.label}</div>
            {#each group.items as item}
                <a class="nav-link {currentPath === item.href ? 'is-active' : ''}" href="{item.href}">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        {@html item.icon}
                    </svg>
                    <span>{item.text}</span>
                </a>
            {/each}
        </nav>
    {/each}

    <div class="sidebar-footer">
        <div class="workspace">
            <div class="workspace-avatar">AD</div>
            <div class="workspace-text">
                <div class="workspace-name">Admin User</div>
                <div class="workspace-role">ADMIN</div>
            </div>
            <svg class="workspace-chev" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                <path d="m7 9 5-5 5 5"/><path d="m7 15 5 5 5-5"/>
            </svg>
        </div>
    </div>
</aside>
