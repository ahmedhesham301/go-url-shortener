import React, { useState } from 'react';

interface ApiResponse {
  id: number;
  long_url: string;
  message: string;
}

const UrlShortener: React.FC = () => {
  const [url, setUrl] = useState('');
  const [originalUrl, setOriginalUrl] = useState('');
  const [shortenedUrl, setShortenedUrl] = useState('');
  const [error, setError] = useState<string>('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost/api/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ url }),
      });

      if (!response.ok) {
        throw new Error('Failed to shorten URL');
      }

      const data: ApiResponse = await response.json();
      setOriginalUrl(data.long_url);
      setShortenedUrl(`${window.location.origin}/${data.id}`);
      setError('');
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred');
      setShortenedUrl('');
      setOriginalUrl('');
    }
  };

  return (
    <div className="url-shortener">
      <h1>URL Shortener</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="url"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          placeholder="Enter a URL to shorten"
          required
        />
        <button type="submit">Shorten URL</button>
      </form>

      {error && <p className="error">{error}</p>}
      
      {shortenedUrl && (
        <div className="result">
          <h2>Shortened URL:</h2>
          {/* <p>Original URL: {originalUrl}</p> */}
          <p>Short URL: <a href={shortenedUrl} target="_blank" rel="noopener noreferrer">
            {shortenedUrl}
          </a></p>
        </div>
      )}
    </div>
  );
};

export default UrlShortener;