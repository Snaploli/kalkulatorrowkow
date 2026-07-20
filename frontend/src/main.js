import './style.css';
import './app.css';

import { CalculatePandQ } from '../wailsjs/go/main/App';

document.addEventListener('DOMContentLoaded', () => {
    const calcBtn = document.getElementById("calculate");
    const externalInput = document.getElementById("external");
    const internalInput = document.getElementById("internal");

    const performCalculation = () => {
        const externalVal = parseFloat(externalInput.value) || 0;
        const internalVal = parseFloat(internalInput.value) || 0;

        const resultContainer = document.getElementById("result");
        if (!resultContainer) return;

        resultContainer.innerHTML = `<div class="result-idle">Obliczanie...</div>`;

        CalculatePandQ(externalVal, internalVal)
            .then((result) => {
                const p = result.p;
                const q = result.q;
                const errMsg = result.errMsg;

                if (errMsg) {
                    resultContainer.innerHTML = `<div class="result-error">Błąd: ${errMsg}</div>`;
                } else {
                    // Display minimalist readout list
                    resultContainer.innerHTML = `
                        <div class="readout-list">
                            <div class="readout-item">
                                <label>Szerokość rowka</label>
                                <span class="readout-value">${q.toFixed(2)}<span class="unit-small">mm</span></span>
                            </div>
                            <div class="readout-item">
                                <label>Średnica podziałowa</label>
                                <span class="readout-value">${p.toFixed(2)}<span class="unit-small">mm</span></span>
                            </div>
                        </div>
                        ${result.suggestedName ? `
                        <div class="suggested-row">
                            <span class="suggested-label">Sugerowany rowek</span>
                            <span class="suggested-value-badge">${result.suggestedName}</span>
                        </div>
                        ` : ''}
                    `;
                }
            })
            .catch((err) => {
                console.error("Error executing Calculate:", err);
                resultContainer.innerHTML = `<div class="result-error">Błąd połączenia</div>`;
            });
    };

    if (calcBtn) {
        // Nasłuchiwanie na kliknięcie przycisku
        calcBtn.addEventListener("click", performCalculation);
    }

    const handleKeydown = (event) => {
        if (event.key === "Enter") {
            performCalculation();
        }
    };

    if (externalInput) {
        externalInput.addEventListener("keydown", handleKeydown);
    }
    if (internalInput) {
        internalInput.addEventListener("keydown", handleKeydown);
    }
});
