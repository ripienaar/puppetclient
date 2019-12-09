// generated code; DO NOT EDIT

package puppetclient

func (c *PuppetClient) debugf(msg string, a ...interface{}) {
	c.clientOpts.logger.Debugf(msg, a...)
}

func (c *PuppetClient) infof(msg string, a ...interface{}) {
	c.clientOpts.logger.Infof(msg, a...)
}

func (c *PuppetClient) warnf(msg string, a ...interface{}) {
	c.clientOpts.logger.Warnf(msg, a...)
}

func (c *PuppetClient) errorf(msg string, a ...interface{}) {
	c.clientOpts.logger.Errorf(msg, a...)
}
