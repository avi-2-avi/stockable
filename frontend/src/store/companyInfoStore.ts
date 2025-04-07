import axios from 'axios';
import { defineStore } from 'pinia'

export const useCompanyInfoStore = defineStore('companyInfo', {
  state: () => ({
    descriptions: {} as Record<string, string>,
    loading: false,
    error: null as string | null,
  }),
  actions: {
    async fetchCompanyDescription(ticker: string, company: string): Promise<string> {
        const key = `${ticker}-${company}`;
        if (this.descriptions[key]) return this.descriptions[key];
        
        const token = localStorage.getItem('auth_token');
        if (!token) {
          throw new Error('No authentication token found');
        }
      
        try {
          const response = await axios.get(`${import.meta.env.VITE_API_URL}/api/company/description?ticker=${ticker}&company=${company}`, {
            headers: {
              'Authorization': `Bearer ${token}`, 
            },
            withCredentials: true, 
          });

          const description = response.data.body.description || 'No info found.';
          this.descriptions[key] = description;
          return description;
        } catch (err) {
          console.error('Error:', err);
          return 'Error fetching company info.';
        }
      },
  },
})
