export function useMask() {
    const formatCurrencyInput = (value: string): string => {
        let cleanValue = value.replace(/\D/g, '');

        cleanValue = cleanValue.replace(/^0+/, '');

        if (!cleanValue) return '0.00';

        if (cleanValue.length > 8) {
            cleanValue = cleanValue.slice(0, 8);
        }

        if (cleanValue.length === 1) {
            return `0.0${cleanValue}`;
        } else if (cleanValue.length === 2) {
            return `0.${cleanValue}`;
        } else {
            const integerPart = cleanValue.slice(0, -2);
            const decimalPart = cleanValue.slice(-2);
            return `${integerPart}.${decimalPart}`;
        }
    };

    const formatCurrencyDisplay = (value: string): string => {
        if (!value) return '';

        const num = Number(value);
        if (isNaN(num)) return '';

        return num.toLocaleString('pt-BR', {
            minimumFractionDigits: 2,
            maximumFractionDigits: 2,
        });
    };

    return {
        formatCurrencyInput,
        formatCurrencyDisplay,
    };
}
