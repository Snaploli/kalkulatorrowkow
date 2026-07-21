import './style.css';
import './app.css';

import { CalculatePandQ } from '../wailsjs/go/main/App';

document.addEventListener('DOMContentLoaded', () => {
    const externalInput = document.getElementById("external");
    const internalInput = document.getElementById("internal");
    const typeSelector = document.getElementById("type-selector");
    const toast = document.getElementById("toast");

    let currentGrooveType = "AUTO";

    // Toast Notification System
    const showToast = (message) => {
        if (!toast) return;
        toast.textContent = message;
        toast.classList.add("show");
        setTimeout(() => {
            toast.classList.remove("show");
        }, 2000);
    };

    const copyToClipboard = (text, label = "") => {
        if (!text) return;
        navigator.clipboard.writeText(text).then(() => {
            showToast(label ? `Skopiowano ${label}: ${text}` : `Skopiowano: ${text}`);
        }).catch((err) => {
            console.error("Clipboard copy failed:", err);
            showToast("Błąd kopiowania do schowka");
        });
    };

    // Auto-select input text on focus for effortless re-entry
    [externalInput, internalInput].forEach(input => {
        if (!input) return;

        input.addEventListener("focus", () => {
            input.select();
        });

        // Live calculation as user types
        input.addEventListener("input", () => {
            performCalculation();
        });
    });

    if (typeSelector) {
        const typeButtons = typeSelector.querySelectorAll(".type-btn");
        typeButtons.forEach(btn => {
            btn.addEventListener("click", () => {
                typeButtons.forEach(b => b.classList.remove("active"));
                btn.classList.add("active");
                currentGrooveType = btn.getAttribute("data-type") || "AUTO";
                performCalculation();
            });
        });
    }

    const performCalculation = () => {
        const externalVal = parseFloat(externalInput.value) || 0;
        const internalVal = parseFloat(internalInput.value) || 0;

        const resultContainer = document.getElementById("result");
        if (!resultContainer) return;

        if (externalVal <= 0 && internalVal <= 0) {
            resultContainer.innerHTML = `<div class="result-idle">Wprowadź wymiary (A) i (B), aby obliczyć rowek.</div>`;
            if (typeSelector && currentGrooveType === "AUTO") {
                const autoBtn = typeSelector.querySelector('[data-type="AUTO"]');
                if (autoBtn) autoBtn.textContent = "Auto";
            }
            return;
        }

        CalculatePandQ(externalVal, internalVal, currentGrooveType)
            .then((result) => {
                const p = result.p;
                const q = result.q;
                const errMsg = result.errMsg;

                // Dynamically update Auto button text to show auto-detected type
                if (typeSelector && currentGrooveType === "AUTO") {
                    const autoBtn = typeSelector.querySelector('[data-type="AUTO"]');
                    if (autoBtn) {
                        autoBtn.textContent = result.grooveType ? `Auto (${result.grooveType})` : "Auto";
                    }
                }

                if (errMsg) {
                    resultContainer.innerHTML = `<div class="result-error">⚠️ ${errMsg}</div>`;
                } else {
                    const isBX = result.grooveType === 'BX';
                    const widthLabel = isBX ? 'Szerokość rowka (N)' : 'Szerokość rowka (F)';
                    const diameterLabel = isBX ? 'Średnica zewnętrzna (∅G)' : 'Średnica podziałowa (∅P)';
                    const diameterValue = isBX ? externalVal : p;

                    const widthValStr = q.toFixed(2);
                    const diaValStr = diameterValue.toFixed(2);
                    const summaryText = `${result.suggestedName ? result.suggestedName + ' | ' : ''}${widthLabel}: ${widthValStr} mm | ${diameterLabel}: ${diaValStr} mm`;

                    resultContainer.innerHTML = `
                        <div class="readout-list">
                            <div class="readout-item clickable" data-copy="${widthValStr}" data-label="szerokość rowka">
                                <div class="readout-label-row">
                                    <label>${widthLabel}</label>
                                    <span class="copy-badge">Kopiuj</span>
                                </div>
                                <span class="readout-value">${widthValStr}<span class="unit-small">mm</span></span>
                            </div>

                            <div class="readout-item clickable" data-copy="${diaValStr}" data-label="średnicę">
                                <div class="readout-label-row">
                                    <label>${diameterLabel}</label>
                                    <span class="copy-badge">Kopiuj</span>
                                </div>
                                <span class="readout-value">${diaValStr}<span class="unit-small">mm</span></span>
                            </div>
                        </div>

                        ${result.suggestedName ? `
                        <div class="suggested-pill clickable" data-copy="${result.suggestedName}" data-label="symbol uszczelki" title="Kliknij, aby skopiować symbol uszczelki">
                            <span class="suggested-pill-label">Sugerowany rowek</span>
                            <span class="suggested-pill-badge">${result.suggestedName}</span>
                        </div>
                        ` : ''}
                    `;

                    // Attach click-to-copy event handlers
                    const copyElements = resultContainer.querySelectorAll('[data-copy]');
                    copyElements.forEach(el => {
                        el.addEventListener("click", (e) => {
                            e.stopPropagation();
                            const valToCopy = el.getAttribute("data-copy");
                            const labelStr = el.getAttribute("data-label") || "";
                            copyToClipboard(valToCopy, labelStr);
                        });
                    });
                }
            })
            .catch((err) => {
                console.error("Error executing Calculate:", err);
                resultContainer.innerHTML = `<div class="result-error">Błąd połączenia z backendem</div>`;
            });
    };

    const handleKeydown = (event) => {
        if (event.key === "Enter") {
            performCalculation();
        } else if (event.key === "Escape") {
            // ESC key clears inputs & resets view
            if (externalInput) externalInput.value = "";
            if (internalInput) internalInput.value = "";
            const resultContainer = document.getElementById("result");
            if (resultContainer) {
                resultContainer.innerHTML = `<div class="result-idle">Wprowadź wymiary (A) i (B), aby obliczyć rowek.</div>`;
            }
            if (externalInput) externalInput.focus();
        }
    };

    if (externalInput) {
        externalInput.addEventListener("keydown", handleKeydown);
    }
    if (internalInput) {
        internalInput.addEventListener("keydown", handleKeydown);
    }

    // Auto-focus first input field on load
    if (externalInput) {
        externalInput.focus();
    }
});
