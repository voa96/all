async function fetchData(url) {
    try {
        const response = await fetch(url);
        if (!response.ok) throw new Error(`HTTP ${response.status}`);
        return await response.json();
    } catch (error) {
        console.error(`Ошибка загрузки ${url}:`, error);
        return null;
    }
}


function formatDate(dateString) {
    const d = new Date(dateString);
    return d.toLocaleString('ru-RU', {
        day: '2-digit',
        month: '2-digit',
        year: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
    });
}

async function loadDashboard() {
    const data = await fetchData('/api/dashboard');
    if (!data) return;

    document.getElementById('totalCameras').textContent = data.total_cameras ?? '—';
    document.getElementById('activeCameras').textContent = data.active_cameras ?? '—';
    document.getElementById('totalEvents').textContent = data.total_events ?? '—';
    document.getElementById('totalArmed').textContent = data.total_armed ?? '—';
    document.getElementById('avgRisk').textContent = data.avg_risk !== undefined ?
        data.avg_risk.toFixed(2) :
        '—';
}


async function loadEvents() {
    const events = await fetchData('/api/events');
    if (!events) return;

    const tbody = document.getElementById('eventsTable');
    document.getElementById('eventsCount').textContent = events.length;

    if (events.length === 0) {
        tbody.innerHTML = `<tr><td colspan="5" style="text-align:center; color:#888; padding:30px;">Нет событий</td></tr>`;
        return;
    }

    tbody.innerHTML = events.map(e => `
        <tr>
            <td>${formatDate(e.timestamp)}</td>
            <td><strong>${e.camera_name || 'Неизвестно'}</strong></td>
            <td>${e.armed_count}</td>
            <td>${e.total_count}</td>
            <td class="risk-${e.risk_level || 'низкий'}">${e.risk_level || '—'}</td>
        </tr>
    `).join('');
}

async function loadCameras() {
    const cameras = await fetchData('/api/cameras');
    if (!cameras) return;

    const tbody = document.getElementById('camerasTable');
    document.getElementById('camerasCount').textContent = cameras.length;

    if (cameras.length === 0) {
        tbody.innerHTML = `<tr><td colspan="3" style="text-align:center; color:#888; padding:30px;">Нет камер</td></tr>`;
        return;
    }

    tbody.innerHTML = cameras.map(c => `
        <tr>
            <td><strong>${c.name}</strong></td>
            <td>${c.location}</td>
            <td class="status-${c.status}">${c.status}</td>
        </tr>
    `).join('');
}


document.addEventListener('DOMContentLoaded', () => {
    loadDashboard();
    loadEvents();
    loadCameras();


    setInterval(() => {
        loadDashboard();
        loadEvents();
    }, 30000);
});

console.log('🚀 Сайт видеонаблюдения загружен!');