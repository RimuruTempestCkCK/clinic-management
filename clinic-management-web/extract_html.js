import fs from 'fs';
import { JSDOM } from 'jsdom';

const html = fs.readFileSync('C:/Users/Maxtop/Downloads/clinic-management/adminator_templete/index.html', 'utf8');

// Load JSDOM and execute scripts
const dom = new JSDOM(html, {
  runScripts: "dangerously",
  resources: "usable",
  url: "file:///C:/Users/Maxtop/Downloads/clinic-management/adminator_templete/index.html"
});

dom.window.addEventListener('load', () => {
  setTimeout(() => {
    const sidebar = dom.window.document.querySelector('[data-shell-sidebar]').innerHTML;
    const topbar = dom.window.document.querySelector('[data-shell-topbar]').innerHTML;
    const footer = dom.window.document.querySelector('[data-shell-footer]').innerHTML;

    fs.writeFileSync('C:/Users/Maxtop/Downloads/clinic-management/clinic-management-web/sidebar_raw.html', sidebar);
    fs.writeFileSync('C:/Users/Maxtop/Downloads/clinic-management/clinic-management-web/topbar_raw.html', topbar);
    fs.writeFileSync('C:/Users/Maxtop/Downloads/clinic-management/clinic-management-web/footer_raw.html', footer);

    console.log("Extraction complete.");
    process.exit(0);
  }, 1000);
});
