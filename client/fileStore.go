package client

func (f *fileStore) Write(text string) {
	f.Mu.Lock()
	defer f.Mu.Unlock()

	f.Store = append(f.Store, text)
}

func (f *fileStore) Read(c *Client) {
	f.Mu.Lock()
	defer f.Mu.Unlock()
	for _, line := range f.Store {
		c.Conn.Write([]byte(line + "\n"))
	}
}
